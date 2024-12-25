package encryption

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"paractice/config"
	"testing"
)

func TestEncryption(t *testing.T) {
	origData := []byte("zse@rgb=")        // 待加密的数据
	key := []byte(config.Config("SALT1")) // 加密的密钥
	log.Println("原文：", string(origData))

	log.Println("------------------ ECB模式 --------------------")
	encrypted := AesEncryptECB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))

	//前端加密后的数据
	decrypted := base64.StdEncoding.EncodeToString(encrypted)
	encrypted, err := base64.StdEncoding.DecodeString(decrypted)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//AES 解密
	//result, _ := AesDecryptECB(encrypted, key)
	//log.Println("解密结果：", string(result))

}
