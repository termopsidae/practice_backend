package test

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"paractice/config"
	"paractice/model/api"
	"paractice/pkg/twitter"
	"strings"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	// 假设前端发送的签名数据
	signature := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
	message := "123456789"                                                               // 假设签名的消息是一串数字
	expectedAddress := common.HexToAddress("0xeA0edc8f212DcB7a39D1cc1106Af4CA4DB0CA545") // 预期的签名者地址

	// 去除签名前缀 "0x"
	signature = strings.TrimPrefix(signature, "0x")

	// 将签名解码为字节数组
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		log.Fatal(err)
	}

	// 使用公钥恢复签名者地址
	recoveredPublicKey, err := crypto.Ecrecover(signatureBytes, signatureBytes)
	if err != nil {
		log.Fatal(err)
	}
	// 从公钥中提取地址
	pk, err := crypto.DecompressPubkey(recoveredPublicKey)
	if err != nil {
		log.Fatal(err)
	}
	recoveredAddress := crypto.PubkeyToAddress(*pk)

	// 验证签名者地址与预期地址是否一致
	if recoveredAddress != expectedAddress {
		fmt.Println("Signature verification failed: recovered address does not match expected address")
		return
	}

	// 打印解析的数字和签名者地址
	fmt.Println("Parsed message:", message)
	fmt.Println("Signer address:", recoveredAddress.Hex())
}
func TestFollow(t *testing.T) {

	err := twitter.Client.FollowUser("PlayStation", "NVBGQ1QtbDBqc09wMGVxQXhHM0dkYmtLUGhHWU1tZUxxMnhtbHlTRmdpWEZCOjE3MTQwMjY1Mzk0Nzk6MToxOmF0OjE")
	fmt.Println(err.Error())
}
func TestPassword(t *testing.T) {
	password := api.Get256Pw("123456")

	fmt.Println(password)
}
func TestTime(t *testing.T) {
	// 获取当前时间并确保是北京时间
	now := time.Now().In(config.BeijingLoc)

	today := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, config.BeijingLoc)
	// 如果当前时间的小时数小于 1 (即凌晨)，则今天是前一天的 0 点
	if now.Hour() < 1 {
		today = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, config.BeijingLoc)
	}
	yesterday := today.AddDate(0, 0, -1)

	// 打印结果，确保是北京时间的 0 点
	fmt.Println("Today: ", today)
	fmt.Println("Yesterday: ", yesterday)
}
