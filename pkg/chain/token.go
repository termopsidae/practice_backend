package chain

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"golang.org/x/net/context"
	"math/big"
	"paractice/pkg"
	"strings"
)

// GetAddressEthBalance 创建交易结构提供前端调用
func GetAddressEthBalance(client *ethclient.Client, addressStr string) (*big.Int, error) {
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(addressStr), nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// GetContractBalanceByAddress 获取指定地址的余额
func GetContractBalanceByAddress(client *ethclient.Client, contractAddress string, address string) (*big.Int, error) {
	hexToAddress := common.HexToAddress(contractAddress)
	contractInstance, err := NewContract(hexToAddress, client)
	if err != nil {
		return nil, err
	}
	return contractInstance.BalanceOf(nil, common.HexToAddress(address))
}

func GetAddrByKey(privateKeyStr string) (*ecdsa.PrivateKey, *common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr) // 不需要0x
	if err != nil {
		return nil, nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, err
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return privateKey, &address, err
}

func EthToWeiByString(ethStr string) (*big.Int, error) {
	decimalNum, err := decimal.NewFromString(ethStr)
	if err != nil {
		return nil, err
	}
	reqParamsNumWei := pkg.EthToWei(&decimalNum) //提现总量 Wei
	return reqParamsNumWei, nil
}

func SendContractBalanceByAddress(client *ethclient.Client,
	contractAddress string,
	formAddressKey *ecdsa.PrivateKey,
	toAddress string,
	toNumWei *big.Int) (*types.Transaction, error) {
	hexToAddress := common.HexToAddress(strings.ToLower(contractAddress))
	contractInstance, err := NewContract(hexToAddress, client)
	if err != nil {
		return nil, err
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
		//log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(formAddressKey, chainID)
	return contractInstance.Transfer(auth, common.HexToAddress(strings.ToLower(toAddress)), toNumWei)
}
