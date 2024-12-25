package pkg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
	"time"
)

// SignatureReceive 领取奖励的Token 签名
func SignatureReceive(contractPrivateKey string, userAddress, tokenAddress string, total, amount, timeUnix *big.Int) (signatureStr string, signatureBytes []byte) {
	// 设置签名的数据
	//id := uint64(1)
	user := common.HexToAddress(userAddress)
	token := common.HexToAddress(tokenAddress)

	// 使用 abi.encodePacked 对数据进行编码
	encodedData := AbiEncodePacked(user, token, total, amount, timeUnix)

	// 创建一个私钥
	privateKey, err := crypto.HexToECDSA(contractPrivateKey)
	if err != nil {
		panic(err)
	}

	// 计算 keccak256 哈希
	hash := crypto.Keccak256(encodedData)
	messageStr := common.Bytes2Hex(hash)
	fmt.Println("消息：", messageStr)

	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		panic(err)
	}
	signatureStr = common.Bytes2Hex(signature)
	fmt.Println("签名结果：", signatureStr)

	// 打印签名结果
	fmt.Println("amount Str:", amount)
	fmt.Println("Message Str:", messageStr)
	fmt.Println("Signature Str:", signatureStr)
	return signatureStr, signature
}

// SignatureIncrease 签名
func SignatureIncrease(contractPrivateKey string, userAddress, tokenAddress string, amount, signatureUnix *big.Int) (signatureStr string, signatureBytes []byte, err error) {
	// 设置签名的数据
	//id := uint64(1)
	user := common.HexToAddress(userAddress)
	token := common.HexToAddress(tokenAddress)
	// 使用 abi.encodePacked 对数据进行编码
	encodedData := AbiEncodePacked(user, token, amount, signatureUnix)

	// 创建一个私钥
	privateKey, err := crypto.HexToECDSA(contractPrivateKey)
	if err != nil {
		return "", nil, err
		//panic(err)
	}

	// 计算 keccak256 哈希
	hash := crypto.Keccak256(encodedData)
	messageStr := common.Bytes2Hex(hash)
	fmt.Println("消息：", messageStr)

	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return "", nil, err
	}
	signatureStr = common.Bytes2Hex(signature)
	fmt.Println("签名结果：", signatureStr)

	// 打印签名结果
	fmt.Println("amount Str:", amount)
	fmt.Println("Message Str:", messageStr)
	fmt.Println("Signature Str:", signatureStr)
	return signatureStr, signature, nil
}

// AbiEncodePacked 使用 abi.encodePacked 对数据进行编码
func AbiEncodePacked(values ...interface{}) []byte {
	var result []byte
	for _, v := range values {
		switch val := v.(type) {
		case common.Address:
			result = append(result, val[:]...)

		case *big.Int:
			result = append(result, common.LeftPadBytes(val.Bytes(), 32)...)
		case int64:
			buf := new(bytes.Buffer)
			binary.Write(buf, binary.BigEndian, val)
			result = append(result, common.LeftPadBytes(buf.Bytes(), 32)...)

		case string:
			result = append(result, []byte(strings.TrimSpace(val))...)
		case time.Time:
			timeInt := big.NewInt(val.UnixNano())
			result = append(result, common.LeftPadBytes(timeInt.Bytes(), 32)...)
		default:
			// 处理其他类型
		}
	}
	return result
}
