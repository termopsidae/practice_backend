package contract

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	chain2 "paractice/chain"
	"paractice/config"
	"paractice/database"
	"paractice/model"
	"paractice/model/api"
	"paractice/pkg"
	"paractice/pkg/chain"
	"strings"

	"testing"
)

func TestCheckTransaction(t *testing.T) {
	ctx := context.Background()
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/") // 更换为你的BSC节点地址
	if err != nil {
		panic(err)
	}

	txHash := common.HexToHash("0x57847963073c036bf618edf6dacbd93632fc9ff68a999d21b2ea99af12c01839") // 替换为你的交易哈希
	_, isPending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		panic(err)
	}
	if isPending {
		fmt.Printf("pending")
	} else {
		fmt.Printf("success")
	}
}
func TestFailNodeMission(t *testing.T) {
	database.ConnectDB()
	chain2.InitBscClient()
	err := api.InitUserTree(database.DB)
	if err != nil {
		panic("init user fail")
	}
	walletAddress := "0x97298fF6eD670870fd2fbC5873dd61b0a1E4874b"
	var nodeId uint
	nodeId = 2
	nodeHash := "0xdf67605992b9f565604ead85633d517cc4a6ecdf53bff600895cfaebe9506c53"
	fmt.Println("/PurchaseNode api...")
	user := model.User{}
	user.WalletAddress = pkg.Upper(walletAddress)
	err = user.GetByWalletAddress(database.DB)
	if err != nil {
		panic("get user fail")
	}
	//-----------------------------验证交易----------------------------
	nodeData, err := model.GetNodeById(database.DB, nodeId)
	if err != nil {
		panic("get node fail")
	}
	toAddr := model.GetByKeyStr(database.DB, "收款地址")
	erc20UsdtAddr := model.GetByKeyStr(database.DB, "BSC的USDT合约地址")
	_, b := chain.CheckTxHashEvent(chain2.BscClient, nodeHash,
		chain.CheckErc20Transfer("", toAddr, erc20UsdtAddr, nodeData.Price))
	if !b {
		panic("get node data fail")
	}
	//认证成功入库成功判断去重复
	err = database.DB.Create(&model.TransactionLog{TxHash: strings.ToLower(nodeHash)}).Error
	if err != nil {
		panic("insert transaction log fail")
	}
	//-----------------------------完成 db入库---------------------------
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		newNodeOrder := model.NodeOrder{
			UserId: user.ID,
			NodeId: nodeId,
			Flag:   "1",
		}
		oid, err := newNodeOrder.InsertNewNodeOrder(tx)
		if err != nil {
			return err
		}
		node := model.Node{}
		node.ID = nodeId
		err = node.GetById(tx)
		if err != nil {
			return err
		}
		if node.Last == "" {
			return errors.New("node has been sold out")
		}
		//赠送矿机
		mining := model.Mining{}
		mining.ID = node.MiningId
		err = mining.GetById(tx)
		if err != nil {
			return err
		}
		newMiningOrder := model.MiningOrder{
			UserId:       user.ID,
			MiningId:     mining.ID,
			ReleaseNum:   mining.AllIncome,
			ReleaseLimit: mining.AllIncome,
			Flag:         "1",
		}
		_, err = newMiningOrder.InsertNewMiningOrder(tx)
		if err != nil {
			return err
		}
		newTransaction := model.Transaction{
			AssociatedId:  oid,
			WalletAddress: user.WalletAddress,
			Hash:          nodeHash,
			Amount:        node.Price,
			ChangeType:    config.TRANSACTION_USDT_PURCHASE,
			Date:          pkg.GetStringToday(),
			Flag:          "1",
		}
		_, err = newTransaction.InsertNewTransaction(tx)
		if err != nil {
			return err
		}
		newTransaction2 := model.Transaction{
			AssociatedId:  oid,
			WalletAddress: user.WalletAddress,
			Amount:        node.BBenefit,
			ChangeType:    config.TRANSACTION_AIGT_AIRDROP,
			Date:          pkg.GetStringToday(),
			Flag:          "1",
		}
		_, err = newTransaction2.InsertNewTransaction(tx)
		if err != nil {
			return err
		}
		newLast := pkg.BigIntStringSub(node.Last, "1")
		nodeParams := map[string]interface{}{"last": newLast}
		err = model.UpdateNode(tx, node.ID, nodeParams)
		if err != nil {
			return err
		}

		newConsume := pkg.BigIntStringAdd(user.Consume, node.Price)
		params := map[string]interface{}{"consume": newConsume}
		if user.Level < node.Level {
			params["level"] = node.Level
		}
		params["b_balance"] = pkg.BigIntStringAdd(user.BBalance, node.BBenefit)
		err = model.UpdateUser(tx, user.ID, params)
		if err != nil {
			return err
		}
		if user.RecommendId != 0 {
			recommendUser := model.User{}
			recommendUser.ID = user.RecommendId
			err = recommendUser.GetById(tx)
			if err != nil {
				return err
			}
			err = UpdateTeamPerformance(tx, recommendUser)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic("transaction fail" + err.Error())
	}

}
