package chain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"paractice/chain"
	"testing"
)

func TestTx(t *testing.T) {
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/") // 更换为你的BSC节点地址
	fmt.Println(client, err)
	_, b := CheckTxHashEvent(client, "0x0e7b3e65e9117511b2b9dcdb7588b97f0d7dbd60fff9a2118bcfa8f4b244a02d",
		CheckErc20Transfer("",
			"0x731286a98c5d80721066981f619a29c764da1091",
			"0x55d398326f99059fF775485246999027B3197955",
			"1000000000000000000000"))
	fmt.Println(b)
}

func TestTx2(t *testing.T) {
	fmt.Println(TopicsErc20Transfer.String())
}

func TestBalance(t *testing.T) {
	addr := "0x2260b9Aa576042aE3b3C9D6Fd565Ed35ae549AC9"
	chain.InitBscClient()
	balanceETH, err := GetAddressEthBalance(chain.BscClient, addr)
	fmt.Println("ETH余额：", balanceETH, err)
	balanceWei, err := GetContractBalanceByAddress(chain.BscClient, "0x55d398326f99059fF775485246999027B3197955", addr)
	fmt.Println("USDT余额：", balanceWei, err)
}

func TestGetAddrByKey(t *testing.T) {
	//对应钱包地址0xc9e1172dEED3e206E988661E7436Bc81ce3603e6  空钱包没用
	keyStr := "22f6c5cbbcf0c0cca79fb0c466046b7d74c0f2060eb0973c7f2cb9b7496415f0"
	key, addr, err := GetAddrByKey(keyStr)
	fmt.Println(key)
	fmt.Println(addr.Hex())
	fmt.Println(err)
}
