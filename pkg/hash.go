package pkg

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"strings"
	"time"
)

func HashStr(sStr string, hStr string) string {
	salt := []byte(sStr) // 自定义盐
	// 对支付密码进行hash并使用自定义盐
	hashedPassword := pbkdf2.Key([]byte(hStr), salt, 4096, 32, sha256.New)
	// 将二进制哈希密码转换为Base64编码，便于存储和传输
	hashedPasswordBase64 := base64.StdEncoding.EncodeToString(hashedPassword)
	return hashedPasswordBase64
}
func SignatureDecode(hash string) (address string, timeStamp time.Time) {
	signature := hash

	// 去除前缀 "0x"
	signature = strings.TrimPrefix(signature, "0x")

	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		log.Fatal(err)
	}
	// 解析签名
	recoveredPublicKey, err := crypto.Ecrecover(common.Hex2Bytes(hash), signatureBytes)
	if err != nil {
		log.Fatal(err)
	}
	// 从公钥中提取地址
	pk, err := crypto.DecompressPubkey(recoveredPublicKey)
	if err != nil {
		log.Fatal(err)
	}
	recoveredAddress := crypto.PubkeyToAddress(*pk)
	fmt.Println("Recovered Address:", recoveredAddress.Hex())
	data := struct {
		Message int64 `json:"message"`
	}{}
	// 定义 EIP-712 数据结构编码器
	//encoder, err := abi.JSON(strings.NewReader(`
	//	{
	//		"message": "int256"
	//	}
	//`))
	//if err != nil {
	//	log.Fatal(err)
	//}
	// 将 Unix 时间戳转换为时间
	timestamp := time.Unix(data.Message, 0)
	return recoveredAddress.String(), timestamp
}
func Upper(data string) string {
	upperStr := strings.ToUpper(data)
	newStr := ""
	for _, char := range upperStr {
		if char >= 'A' && char <= 'Z' {
			newStr += string(char)
		} else {
			newStr += strings.ToUpper(string(char))
		}
	}
	return newStr
}
func IsValidAddress(data string) bool {
	if common.IsHexAddress(data) {
		return true
	}
	return false
}
