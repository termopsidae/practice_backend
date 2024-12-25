package chain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"paractice/config"
	"strings"
)

type Transaction struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      uint64 `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Value    string `json:"value"`
	Data     string `json:"data"`
	Nonce    uint64 `json:"nonce"`
}

var parsedABI, _ = abi.JSON(strings.NewReader(`[{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"constant":true,"inputs":[],"name":"_decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"_name","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"_symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"burn","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getOwner","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"mint","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"renounceOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"recipient","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"sender","type":"address"},{"internalType":"address","name":"recipient","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))

var contractABI, _ = abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"address","name":"_usdt","type":"address"},{"internalType":"address","name":"_usdc","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"USDC","outputs":[{"internalType":"contract IERC20","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"USDT","outputs":[{"internalType":"contract IERC20","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"adminAddress","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getAdmin","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getSellData","outputs":[{"components":[{"internalType":"address","name":"add","type":"address"},{"internalType":"uint256","name":"num","type":"uint256"}],"internalType":"struct Data.SellData","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_tokenAddr","type":"address"}],"name":"getTokenBalance","outputs":[{"internalType":"uint256","name":"tokenBalance","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"}],"name":"getUSDCPrice","outputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"}],"name":"gutUsdtPrice","outputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"router","outputs":[{"internalType":"contract IPancakeRouter02","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"}],"name":"sell","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"sellIndex","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newAdmin","type":"address"}],"name":"setAdmin","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_USDCAddr","type":"address"}],"name":"setUSDCAddress","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_USDTAddr","type":"address"}],"name":"setUSDTAddress","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"tokenAddress","type":"address"},{"internalType":"address","name":"toAddress","type":"address"},{"internalType":"uint256","name":"num","type":"uint256"}],"name":"wToken","outputs":[],"stateMutability":"nonpayable","type":"function"}]`))

// CreateErc20Transfer 创建交易结构提供前端调用
func CreateErc20Transfer(c *ethclient.Client, fromAddressStr, toAddressStr, erc20AddressStr, amountWeiStr string) (*Transaction, error) {
	fromAddr := common.HexToAddress(strings.ToLower(fromAddressStr))                   //发送方
	toAddress := common.HexToAddress(strings.ToLower(toAddressStr))                    //代币接收方
	erc20TokenContractAddress := common.HexToAddress(strings.ToLower(erc20AddressStr)) //代币合约

	amountWei := new(big.Int)
	amountWei.SetString(amountWeiStr, 10) // 10 USDT
	data, err := parsedABI.Pack("transfer", toAddress, amountWei)
	if err != nil {
		return nil, err
	}
	gasLimit := uint64(300000)
	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	nonce, err := c.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		return nil, err
	}
	return &Transaction{
		From:     fromAddr.Hex(),
		To:       erc20TokenContractAddress.Hex(),
		Gas:      gasLimit,
		GasPrice: gasPrice.String(),
		Data:     fmt.Sprintf("0x%x", data),
		Value:    "0x0",
		Nonce:    nonce,
	}, nil
}

func FindUserAddressTokenBalance(client *ethclient.Client, token, address string) (*big.Int, error) {
	erc20ContractAddress := common.HexToAddress(strings.ToLower(token)) //代币合约
	contractNew, err := NewContract(erc20ContractAddress, client)
	if err != nil {
		return nil, err
	}
	callOpts := &bind.CallOpts{Pending: false, Context: context.Background()}
	result, err := contractNew.BalanceOf(callOpts, common.HexToAddress(strings.ToLower(address)))
	return result, nil
}

func FindUserAddressTokenApproveAmount(client *ethclient.Client, token, userAddress, toAddress string) (*big.Int, error) {
	erc20ContractAddress := common.HexToAddress(strings.ToLower(token)) //代币合约
	contractNew, err := NewContract(erc20ContractAddress, client)
	if err != nil {
		return nil, err
	}
	callOpts := &bind.CallOpts{Pending: false, Context: context.Background()}
	result, err := contractNew.Allowance(callOpts, common.HexToAddress(strings.ToLower(userAddress)), common.HexToAddress(strings.ToLower(toAddress)))
	return result, nil
}

func ApproveToken(client *ethclient.Client, tokenAddressStr, fromAddressStr, toAddressStr string, amountWei *big.Int) (*types.Transaction, error) {
	ctx := context.Background()
	fromAddr := common.HexToAddress(strings.ToLower(fromAddressStr))              //发送方
	toAddress := common.HexToAddress(strings.ToLower(toAddressStr))               //代币接收方
	tokenContractAddress := common.HexToAddress(strings.ToLower(tokenAddressStr)) //代币合约

	data, err := parsedABI.Pack("approve", toAddress, amountWei)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	gasLimit := uint64(300000)

	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return nil, err
	}
	baseTx := &types.LegacyTx{
		To:       &tokenContractAddress,
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		Value:    new(big.Int),
		Data:     data,
	}
	return types.NewTx(baseTx), nil
}

func ConvertTransactionResp(fromAddress string, info *types.Transaction) Transaction {
	value := "0x0"
	if info.Value() != nil {
		value = fmt.Sprintf("0x%v", info.Value().String())
	}
	fmt.Println("data", info.Data())
	return Transaction{
		From:     fromAddress,
		To:       info.To().String(),
		Gas:      info.Gas(),
		GasPrice: info.GasPrice().String(),
		Data:     fmt.Sprintf("0x%x", info.Data()),
		Value:    value,
		Nonce:    info.Nonce(),
	}
}

// RechargeUSDT 充值USDT ，拿USDC 卖USDT
func RechargeUSDT(c *ethclient.Client, fromAddress string, usdtValue *big.Int) (*types.Transaction, error) {
	fromAddr := common.HexToAddress(strings.ToLower(fromAddress))                                           //发送方
	contractAddress := common.HexToAddress(strings.ToLower(config.Config("BSC_RECHARGE_CONTRACT_ADDRESS"))) //代币合约

	data, err := contractABI.Pack("sell", usdtValue)
	if err != nil {
		return nil, err
	}
	gasLimit := uint64(300000)
	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	nonce, err := c.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		return nil, err
	}
	baseTx := &types.LegacyTx{
		To:       &contractAddress,
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		Value:    new(big.Int),
		Data:     data,
	}
	return types.NewTx(baseTx), nil
}

// RechargeContractWithdraw 充值合约 提现的
func RechargeContractWithdraw(c *ethclient.Client, tokenAddress, toAddress string, usdcValue *big.Int) (*types.Transaction, error) {
	fromAddr := common.HexToAddress(strings.ToLower(config.Config("RECHARGE_CONTRACT_ADMIN_ADDRESS")))      //发送方
	tokenAddr := common.HexToAddress(strings.ToLower(tokenAddress))                                         //代币地址
	toAddr := common.HexToAddress(strings.ToLower(toAddress))                                               //接收方
	contractAddress := common.HexToAddress(strings.ToLower(config.Config("BSC_RECHARGE_CONTRACT_ADDRESS"))) //代币合约

	data, err := contractABI.Pack("wToken", tokenAddr, toAddr, usdcValue)
	if err != nil {
		return nil, err
	}
	gasLimit := uint64(300000)
	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	nonce, err := c.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		return nil, err
	}
	baseTx := &types.LegacyTx{
		To:       &contractAddress,
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		Value:    new(big.Int),
		Data:     data,
	}
	return types.NewTx(baseTx), nil
}
