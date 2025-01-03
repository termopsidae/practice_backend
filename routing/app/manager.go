package app

import (
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"paractice/config"
	"paractice/database"
	"paractice/model"
	"paractice/pkg"
	"paractice/pkg/encryption"
	"paractice/routing/types"
	"strconv"
	"strings"
	"time"
)

// Login 登录
func Login(c *fiber.Ctx) error {
	fmt.Println("/Login api...")
	reqParams := types.LoginReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", config.MESSAGE_PARSER_ERROR))
	}

	inputPassword := reqParams.Password
	if inputPassword == "" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", config.MESSAGE_PARSER_ERROR))
	}
	encrypted, err := base64.StdEncoding.DecodeString(inputPassword)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", config.MESSAGE_PARSER_ERROR))
	}
	//AES 解密
	//result, err := encryption.AesDecryptECB(encrypted, []byte(config.Config("SALT1")))
	result, err := encryption.AesDecryptECB(encrypted, []byte(config.Config("SALT1")))
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", "解密错误"))
	}
	password := string(result)
	manager := model.Manager{}
	manager.UserName = reqParams.ManagerName
	err = manager.GetByUserName(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get manager error", ""))
	}
	returnT := ""
	data := types.LoginResp{
		Class: manager.Class,
	}
	if password != manager.Password {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "wrong password", "密码错误"))
	}
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		if password == manager.Password {
			returnT = strings.Split(manager.Token, ":")[0]
			c.Locals(config.LOCAL_TOKEN, returnT)
			data.Token = returnT
		}
		if !pkg.CheckTokenValidityTime(&manager.Token) {
			returnT = pkg.RandomString(64)
			params := map[string]interface{}{}
			params["token"] = returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10)
			err = model.UpdateManager(database.DB, manager.ID, params)
			if err != nil {
				return err
			}
			c.Locals(config.LOCAL_TOKEN, returnT)
			data.Token = returnT
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
	}
	c.Locals(config.LOCAL_TOKEN, returnT)
	return c.JSON(pkg.SuccessResponse(data))
}

// 创建商品,
func CreateNewGood(c *fiber.Ctx) error {
	reqParams := types.CreatNewGoodReq{}
	err := c.BodyParser(&reqParams) //获得传进来的请求参数名称
	if err != nil {                 // 说明传入参数格式不正确
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", config.MESSAGE_PARSER_ERROR))
	}
	if reqParams.Price == 0 || reqParams.GoodName == "" || reqParams.Description == "" || reqParams.LastAmount == 0 {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", config.MESSAGE_PARSER_ERROR))
	}
	//获取获取当前登录管理人员

	managerId := c.Locals(config.LOCAL_MANAGERID_INT64).(int64)
	currentManager, err := model.SelectManagerById(database.DB, uint(managerId))
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "services exception", config.MESSAGE_GET_TRANSACTION_ERROR))
	}
	//对管理员是否有操作商品权限进行判断 Class 为 "1" 表示有权限  "2" 表示无权限
	if currentManager.Class != "1" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "no permissions", "你没有操作权限!"))
	}
	good := model.Good{
		Name:        reqParams.GoodName,
		Price:       reqParams.Price,
		LastAmount:  reqParams.LastAmount,
		Description: reqParams.Description,
		Flag:        "1",
	}
	transactionError := database.DB.Transaction(func(tx *gorm.DB) error {
		good.InsertNewGood(tx) //插入新商品
		return nil
	})
	if transactionError != nil { //说明添加商品失败，
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "add good fail, please try again", "添加商品失败，请重试"))
	}
	// 如果没有拦截 成功返回
	repParams := types.CreatNewGoodResp{}
	return c.JSON(pkg.SuccessResponse(repParams))
}

// 根据条件查询商品
func SelectGoodList(c *fiber.Ctx) error {
	reqParams := types.SelectGoodListReq{}
	err := c.BodyParser(&reqParams)
	//if err != nil {
	//	return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", config.MESSAGE_PARSER_ERROR))
	//}
	managerId := c.Locals(config.LOCAL_MANAGERID_INT64).(int64)
	currentManager, err := model.SelectManagerById(database.DB, uint(managerId))
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "services exception", config.MESSAGE_GET_TRANSACTION_ERROR))
	}
	//对管理员是否有操作商品权限进行判断 Class 为 "1" 表示有权限  "2" 表示无权限
	if currentManager.Class != "1" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "no permissions", "你没有操作权限!"))
	}
	//根据参数 查出商品 并更具时间排序
	goods, err := model.SelectGoodsByCondition(database.DB, reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "select fail", config.MESSAGE_GET_TRANSACTION_ERROR))
	}
	//开始组装数据
	typeGoods := make([]types.TypeGood, 0)
	for _, good := range goods {
		typeGood := types.TypeGood{
			GoodName:    good.Name,
			Price:       good.Price,
			LastAmount:  good.LastAmount,
			Description: good.Description,
			Flag:        good.Flag,
		}
		typeGoods = append(typeGoods, typeGood)
	}
	//开始组装返回数据
	repData := types.SelectGoodListResp{
		GoodsList: typeGoods,
	}
	//返回成功数据
	return c.JSON(pkg.SuccessResponse(repData))
}
