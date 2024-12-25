package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"
)

var TopicsErc20Transfer = common.BytesToHash(crypto.Keccak256([]byte("Transfer(address,address,uint256)")))

func CheckTxHashEvent(client *ethclient.Client, txHashStr string, checkFunc func(log *types.Log) bool) (*types.Log, bool) {
	if len(txHashStr) != 66 {
		return nil, false
	}
	txHash := common.HexToHash(txHashStr)
	var receipt *types.Receipt
	var err error
	for i := 0; i < 4; i++ {
		receipt, err = client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			break
		}
		time.Sleep(time.Second * 5)
	}
	if err != nil {
		return nil, false
	}
	for _, log := range receipt.Logs {
		if checkFunc(log) {
			return log, true
		}
	}
	return nil, false
}

type TransferData1 struct {
	From       common.Address
	To         common.Address
	Erc20Token common.Address
	Num        big.Int
}

func LogToTransferData(log *types.Log) *TransferData1 {
	if log == nil {
		return nil
	}
	if len(log.Topics) != 3 {
		return nil
	}
	if log.Topics[0] != TopicsErc20Transfer { //判断是否是转账事件
		return nil
	}
	from := common.BytesToAddress(log.Topics[1].Bytes())
	to := common.BytesToAddress(log.Topics[2].Bytes())
	logValue := new(big.Int).SetBytes(log.Data)
	return &TransferData1{
		From:       from,
		To:         to,
		Erc20Token: log.Address,
		Num:        *logValue,
	}
}

/*
CheckErc20Transfer 检查ERC20 转账

	fromAddress=发送方 传“” 则不校验
	toAddress=接收方 传“” 则不校验
	erc20Address=代币地址
	value=金额要求wei但闻  传“” 则不校验
*/
func CheckErc20Transfer(fromAddress, toAddress, erc20Address, value string) func(log *types.Log) bool {
	return func(log *types.Log) bool {
		if log == nil {
			return false
		}
		if len(log.Topics) != 3 {
			return false
		}
		if log.Topics[0] != TopicsErc20Transfer { //判断是否是转账事件
			return false
		}
		if strings.ToLower(log.Address.String()) != strings.ToLower(erc20Address) { //判断那个代币合约事件
			return false
		}

		//校验发送方
		if fromAddress != "" {
			from := common.BytesToAddress(log.Topics[1].Bytes())
			if strings.ToLower(fromAddress) != strings.ToLower(from.String()) {
				return false
			}
		}

		//校验接收方
		if toAddress != "" {
			to := common.BytesToAddress(log.Topics[2].Bytes())
			if strings.ToLower(toAddress) != strings.ToLower(to.String()) {
				return false
			}
		}
		logValue := new(big.Int).SetBytes(log.Data)
		if value != "" {
			intValue, b := new(big.Int).SetString(value, 10)
			if !b {
				return false
			}
			//	-1 if x <  y
			//	 0 if x == y
			//	+1 if x >  y
			i := logValue.Cmp(intValue)
			if i == -1 { //金额不对
				return false
			}
		}
		//不限制金额
		if value == "" {
			intValue := new(big.Int).SetInt64(0)
			i := logValue.Cmp(intValue)
			if i == -1 || i == 0 { //金额不对
				return false
			}
		}
		return true
	}
}

/*
CheckErc20TransferByErc20Arr 检查ERC20 转账

	fromAddress=发送方 传“” 则不校验
	toAddress=接收方 传“” 则不校验
	erc20Address= 代币地址 数组
	value=金额要求wei但闻  传“” 则不校验
*/
func CheckErc20TransferByErc20Arr(fromAddress, toAddress string, erc20AddressArr []string, value string) func(log *types.Log) bool {
	return func(log *types.Log) bool {
		if log == nil {
			return false
		}
		if len(log.Topics) != 3 {
			return false
		}
		if log.Topics[0] != TopicsErc20Transfer { //判断是否是转账事件
			return false
		}

		//检查是否包含制定代币转账
		erc20AddressBool := false
		for _, erc20Address := range erc20AddressArr {
			if strings.ToLower(log.Address.String()) == strings.ToLower(erc20Address) { //判断那个代币合约事件
				erc20AddressBool = true
			}
		}

		if !erc20AddressBool {
			return false
		}

		//校验发送方
		if fromAddress != "" {
			from := common.BytesToAddress(log.Topics[1].Bytes())
			if strings.ToLower(fromAddress) != strings.ToLower(from.String()) {
				return false
			}
		}

		//校验接收方
		if toAddress != "" {
			to := common.BytesToAddress(log.Topics[2].Bytes())
			if strings.ToLower(toAddress) != strings.ToLower(to.String()) {
				return false
			}
		}
		logValue := new(big.Int).SetBytes(log.Data)
		if value != "" {
			intValue, b := new(big.Int).SetString(value, 10)
			if !b {
				return false
			}
			//	-1 if x <  y
			//	 0 if x == y
			//	+1 if x >  y
			i := logValue.Cmp(intValue)
			if i == -1 { //金额不对
				return false
			}
		}
		//不限制金额
		if value == "" {
			intValue := new(big.Int).SetInt64(0)
			i := logValue.Cmp(intValue)
			if i == -1 || i == 0 { //金额不对
				return false
			}
		}
		return true
	}
}
