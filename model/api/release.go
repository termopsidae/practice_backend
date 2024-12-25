package api

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"paractice/config"
)

func ReleaseERC20(to_address string, transfer_amount *big.Int, token_type string) (string, error) {
	rpcURL := config.Config("CHAIN_RPC_URL")
	// 转账的私钥
	privateKey := config.Config("CONTRACT_PRIVATE_KEY")
	var tokenAddress string
	if token_type == "usdt" {
		tokenAddress = config.Config("USDT_ADDRESS_TEST")
	} else if token_type == "unc" {
		tokenAddress = config.Config("UNC_ADDRESS_TEST")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", err
	}
	// 解析私钥
	privateKeyBytes, err := hexutil.Decode(privateKey)
	if err != nil {
		return "", err
	}
	key, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return "", err
	}
	// 创建交易
	auth := bind.NewKeyedTransactor(key)
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // 转账的ETH数量，如果是转账ETH则需要设置
	auth.GasLimit = uint64(1000000) // 转账的Gas上限
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	//address := common.HexToAddress(tokenAddress)
	//instance, err := contract.NewContracts(address, client)
	//if err != nil {
	//	return "", err
	//}
	//to := common.HexToAddress(to_address)
	//tx, err := instance.Transfer(auth, to, transfer_amount)
	//if err != nil {
	//	return "", err
	//}
	//fmt.Println("Transferring " + token_type + "...")
	//// 等待交易被打包
	//receipt, err := bind.WaitMined(context.Background(), client, tx)
	//if err != nil {
	//	return "", err
	//}
	return tokenAddress, nil
}
