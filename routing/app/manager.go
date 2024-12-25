package app

import (
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"paractice/config"
	"paractice/database"
	"paractice/model"
	"paractice/pkg"
	"paractice/pkg/encryption"
	"paractice/routing/types"
	"strconv"
	"strings"
	"time"
)

// Login 登录
func Login(c *fiber.Ctx) error {
	fmt.Println("/Login api...")
	reqParams := types.LoginReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", config.MESSAGE_PARSER_ERROR))
	}

	inputPassword := reqParams.Password
	if inputPassword == "" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", config.MESSAGE_PARSER_ERROR))
	}
	encrypted, err := base64.StdEncoding.DecodeString(inputPassword)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", config.MESSAGE_PARSER_ERROR))
	}
	//AES 解密
	//result, err := encryption.AesDecryptECB(encrypted, []byte(config.Config("SALT1")))
	result, err := encryption.AesDecryptECB(encrypted, []byte(config.Config("SALT1")))
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", "解密错误"))
	}
	password := string(result)
	manager := model.Manager{}
	manager.UserName = reqParams.ManagerName
	err = manager.GetByUserName(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get manager error", ""))
	}
	returnT := ""
	data := types.LoginResp{
		Class: manager.Class,
	}
	if password != manager.Password {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "wrong password", "密码错误"))
	}
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		if password == manager.Password {
			returnT = strings.Split(manager.Token, ":")[0]
			c.Locals(config.LOCAL_TOKEN, returnT)
			data.Token = returnT
		}
		if !pkg.CheckTokenValidityTime(&manager.Token) {
			returnT = pkg.RandomString(64)
			params := map[string]interface{}{}
			params["token"] = returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10)
			err = model.UpdateManager(database.DB, manager.ID, params)
			if err != nil {
				return err
			}
			c.Locals(config.LOCAL_TOKEN, returnT)
			data.Token = returnT
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), config.MESSAGE_TRANSACTION_ERROR))
	}
	c.Locals(config.LOCAL_TOKEN, returnT)
	return c.JSON(pkg.SuccessResponse(data))
}
