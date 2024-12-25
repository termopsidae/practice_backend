package test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"paractice/pkg"
	"testing"
)

func TestRsa1(t *testing.T) {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// 明文消息
	message := "Hello, world!"

	// 加密明文消息
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		[]byte(message),
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后的消息：%x\n", ciphertext)

	// 解密密文消息
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		ciphertext,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后的消息：%s\n", plaintext)
}

func TestRsa2(t *testing.T) {
	hash := pkg.RSAEncrypt("dadssda", "")
	fmt.Println(hash)

	fmt.Println(pkg.RSADecrypt(hash, ""))

}
func TestRsa33(t *testing.T) {
	//pkg.RSADecrypt("Hello, world!")
}
func TestApprove(t *testing.T) {
	//contract.OWNER_PK = "f87db617972f7d5f50441a8ee919583d2f4da9cf16a5533c10dd20a3b50e5372"
	//contract.ACCOUNT_PK = "b64b03e89d61a76b2f4a780c958698f27351a6b74424be3d030ab86da5af2e26"
	//tokenBalance, err := contract.GetTokenBalance("0xb6AAb4C32Af16DF75D215b62f9CeD43176A083F8")
	//if err != nil {
	//	println(tokenBalance.String())
	//}
	//var a big.Int
	//a.SetString(tokenBalance.String(), 10)
	//contract.ApproveToken(&a, "0xb6AAb4C32Af16DF75D215b62f9CeD43176A083F8")
	//balance, err := contract.AllowanceToken("0xb6AAb4C32Af16DF75D215b62f9CeD43176A083F8")
	//if err != nil {
	//	return
	//}
	//println(balance.String())

}
