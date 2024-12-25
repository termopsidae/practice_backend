package pkg

import (
	"errors"
	"github.com/pquerna/otp/totp"
	"time"
)

func GenSecretKey(phone string) string {
	// 生成一个新的秘密密钥
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Pooluo",
		AccountName: "phone",
	})
	if err != nil {
		panic(err)
	}
	ks := key.Secret()
	return ks

}

func GenCode(ks string) string {
	// 创建一个新的TOTP生成器
	code, err := totp.GenerateCode(ks, time.Now())
	if err != nil {
		panic(err)
	}
	return code
}

func ValidCode(code, ks string) error {
	// 验证一个验证码是否有效
	isValid := totp.Validate(code, ks)
	if isValid {
		return nil
	} else {
		return errors.New("验证码无效")
	}
}
