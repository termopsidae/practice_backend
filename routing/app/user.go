package app

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"paractice/config"
	"paractice/database"
	"paractice/model"
	"paractice/model/api"
	"paractice/pkg"
	"paractice/routing/types"
	"strconv"
	"strings"
	"time"
)

// RegisterAndLogin 登录注册
func RegisterAndLogin(c *fiber.Ctx) error {
	fmt.Println("/Register api...")
	reqParams := types.RegisterAndLoginReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", config.MESSAGE_PARSER_ERROR))
	}
	if !pkg.IsValidAddress(reqParams.WalletAddress) {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "invalid address error", config.MESSAGE_ADDRESS_ERROR))
	}
	signature, err := hex.DecodeString(strings.TrimPrefix(reqParams.Signature, "0x"))
	if err != nil || signature == nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "wrong signature", "签名信息错误 001"))
	}
	address, err := ecRecover(ethMessage(reqParams.Message), signature)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "wrong signature", "签名信息错误 011"))
	}
	//签名失败
	if strings.ToLower(address.Hex()) != reqParams.WalletAddress {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "wrong signature", "签名信息错误 002"))
	}
	returnT := pkg.RandomString(64)
	data := types.RegisterAndLoginResp{}
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		user := model.User{WalletAddress: pkg.Upper(reqParams.WalletAddress)}
		err = user.GetByWalletAddress(database.DB)
		if err != nil {
			if !strings.Contains(err.Error(), "record not found") {
				return err
			}
			newUser := model.User{
				WalletAddress: pkg.Upper(reqParams.WalletAddress),

				Token: returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10),
				Flag:  "0",
			}

			id, err := newUser.InsertNewUser(tx)
			if err != nil {
				return err
			}
			lq := make(map[time.Time]string)
			api.LockedMap[id] = api.LockedQueue{AllAmount: "0", SingleLocked: lq}
			api.LockedAMap[id] = api.LockedQueue{AllAmount: "0", SingleLocked: lq}
			api.LockedBMap[id] = api.LockedQueue{AllAmount: "0", SingleLocked: lq}
			data.Token = returnT
		} else {
			if !pkg.CheckTokenValidityTime(&user.Token) {
				params := map[string]interface{}{
					"token": returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10),
				}
				err = model.UpdateUserByAddress(tx, pkg.Upper(reqParams.WalletAddress), params)
				if err != nil {
					return err
				}
				data.Token = returnT
			}
			data.Token = strings.Split(user.Token, ":")[0]
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
	}
	c.Locals(config.LOCAL_TOKEN, returnT)
	return c.JSON(pkg.SuccessResponse(data))
}

// SelectAllGoods 查询所有商品接口
func SelectAllGoods(c *fiber.Ctx) error {
	data := types.SelectAllGoodsResp{}
	//err = database.DB.Transaction(func(tx *gorm.DB) error {
	//
	//})
	//查询数据,不需要包裹数据库事物
	list := make([]types.GoodInfo, 0)
	goods, err := model.SelectAllGoods(database.DB, "1")
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
	} else {
		for _, g := range goods {
			in := types.GoodInfo{
				GoodId:      g.ID,
				Name:        g.Name,
				Price:       g.Price,
				Description: g.Description,
				LastAmount:  g.LastAmount,
				Flag:        g.Flag,
			}
			list = append(list, in)
		}
	}
	data.Goods = list
	return c.JSON(pkg.SuccessResponse(data))
}

// PurchaseGood 添加商品到购物车
func PurchaseGood(c *fiber.Ctx) error {
	reqParams := types.PurchaseGoodReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", config.MESSAGE_PARSER_ERROR))
	}
	//通过 good id ，amount数量 创建 订单

	currentUser := c.Locals(config.LOCAL_USERID_STRUCT).(model.User)

	//3，通过用户id 反查 匹配的购物车
	cart, err := model.SelectCartByUserId(database.DB, currentUser.ID)
	if err != nil { //没有购物车，为当前用户插入购物车
		if strings.Contains(err.Error(), "record not found") {
			cart = model.Cart{
				UserId:          currentUser.ID,
				User:            currentUser,
				GoodOrderIdList: "[]",
				TotalPrice:      0.0,
				Flag:            "1",
			}
			transactionErr := database.DB.Transaction(func(tx *gorm.DB) error {
				//订单加入购物车
				cartId, err := cart.InsertCart(database.DB)
				if err != nil {
					return err
				}
				cart.ID = cartId
				return nil
			})
			if transactionErr != nil {
				return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, transactionErr.Error(), config.MESSAGE_TRANSACTION_ERROR))
			}
		} else {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
		}
	}
	//根据用户传入参数 创建订单
	good, err := model.SelectGoodById(database.DB, reqParams.GoodId)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "haven't this good", "没有此商品"))
	}
	amount := reqParams.Amount //购买数量
	price := good.Price        //商品单价
	totalPrice := float64(amount) * price
	//创建订单
	goodOrder := model.GoodOrder{
		UserId:     currentUser.ID,
		GoodId:     good.ID,
		Good:       good,
		CartId:     cart.ID,
		Amount:     amount,
		TotalPrice: totalPrice,
		Flag:       "1",
	}
	id, orderErr := goodOrder.InsertNewGoodOrder(database.DB)
	if orderErr != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
	}
	goodOrderList := make([]uint, 0)
	err = json.Unmarshal([]byte(cart.GoodOrderIdList), &goodOrderList)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
	}
	goodOrderList = append(goodOrderList, id)
	goodOrderListStr, err := json.Marshal(goodOrderList)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
	}

	transactionErr := database.DB.Transaction(func(tx *gorm.DB) error {
		//更新购物车
		params := map[string]interface{}{
			"total_price":        cart.TotalPrice + goodOrder.TotalPrice,
			"good_order_id_list": string(goodOrderListStr),
		}
		err = model.UpdateCart(tx, cart.ID, params)
		if err != nil {
			return err
		}
		return nil
	})
	if transactionErr != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
	}
	//返回成功给前端
	responseData := types.PurchaseGoodResp{}
	return c.JSON(pkg.SuccessResponse(responseData))
}

// SelectCartInfo 查询购物车,仅仅查询待付款
func SelectCartInfo(c *fiber.Ctx) error {
	//reqParams := types.SelectCartInfoReq{}
	//err := c.BodyParser(&reqParams)
	//if err != nil {
	//	return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "passer error", config.MESSAGE_PARSER_ERROR))
	//}
	currentUser := c.Locals(config.LOCAL_USERID_STRUCT).(model.User)
	//查询出当前用户的购物车
	cart, err := model.SelectCartByUserId(database.DB, currentUser.ID)
	if err != nil { //没有购物车 为当前用户插入购物车
		if strings.Contains(err.Error(), "record not found") { //查询购物车无记录条件
			//此时我们应为当前用户创建个购物车并插入数据库
			cart := model.Cart{
				UserId:          currentUser.ID,
				User:            currentUser,
				GoodOrderIdList: "[]",
				TotalPrice:      0.0,
				Flag:            "1",
			}
			transactionErr := database.DB.Transaction(func(tx *gorm.DB) error {
				cartId, err := cart.InsertCart(tx)
				if err != nil {
					return err
				}
				cart.ID = cartId
				return nil
			})
			if transactionErr != nil { //如果插入失败，就报服务异常
				return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, transactionErr.Error(), config.MESSAGE_TRANSACTION_ERROR))
			}
		} else { //查询购物车记录错误
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
		}
	}
	goodOrderIdList := make([]uint, 0)
	err = json.Unmarshal([]byte(cart.GoodOrderIdList), &goodOrderIdList)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
	}
	//查询数据库中购物车中所有还没购买的商品
	goodOrderList, err := model.SelectOrderByIdAndFlag(database.DB, goodOrderIdList, "1")
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
	}
	goodOrderInfoList := make([]types.GoodOrderInfo, 0)
	//将从数据库查询出来的购物车订单数据复制给 types 中 GoodOrderInfo 类
	for _, goodOrder := range goodOrderList {
		//通过当前订单的商品Id 查询出商品详情
		good, err := model.SelectGoodById(database.DB, goodOrder.GoodId)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
		}
		goodOrderInfo := types.GoodOrderInfo{
			GoodOrderId: goodOrder.ID,
			Good: types.GoodInfo{
				GoodId:      goodOrder.GoodId,
				Name:        good.Name,
				Price:       good.Price,
				Description: good.Description,
				LastAmount:  good.LastAmount,
				Flag:        good.Flag,
			},
			Amount:     goodOrder.Amount,
			TotalPrice: goodOrder.TotalPrice,
		}
		goodOrderInfoList = append(goodOrderInfoList, goodOrderInfo)
	}
	//开始组装返回结果数据
	resPonseData := types.SelectCartInfoResp{
		CartId:        cart.ID,
		TotalPrice:    cart.TotalPrice,
		GoodOrderList: goodOrderInfoList,
	}
	//数据组装成功,以成功格式返回
	return c.JSON(pkg.SuccessResponse(resPonseData))
}

// 解签名
func ecRecover(sighash []byte, sig []byte) (common.Address, error) {
	if len(sig) < 64 {
		return [20]byte{}, errors.New("err")
	}
	sig[64] -= 27
	defer func() { sig[64] += 27 }()

	signer, err := crypto.SigToPub(sighash, sig)
	if err != nil {
		//utils.Fatalf("Failed to recover sender from signature %x: %v", sig, err)
		return [20]byte{}, err
	}
	return crypto.PubkeyToAddress(*signer), nil
}

// keccak256签名
func ethMessage(message string) []byte {
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	return crypto.Keccak256([]byte(prefix))
}
