package intcpt

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"math/big"
	"paractice/config"
	"paractice/database"
	"paractice/model"
	"paractice/pkg"
)

// AuthApp Protected protect routes
func AuthApp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//var token = c.Get("token")
		//fmt.Println(token)
		//c.Locals("guo", name)
		//c.Locals("guo2", "23123123123123123")
		var (
			id    uint
			token = c.Get(config.LOCAL_TOKEN)
			db    = database.DB
			err   error
		)

		//开关 测试打开token
		//if true {
		//	c.Locals(config.LOCAL_USERID_UINT, 1)
		//	c.Locals(config.LOCAL_USERID_INT64, 1)
		//	_ = c.Next()
		//	return nil
		//}

		// 打印请求地址
		log.Info("Request URL: ", c.Path())
		log.Info("Request JSON: ", string(c.Body()))
		if token == "" || len(token) < 10 {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is null", ""))
		}

		id, tokenData, err := model.UserSelectAddressByToken(db, token)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		user := model.User{}
		user.ID = id
		err = user.GetById(db)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		if !pkg.CheckSpecialCharacters(&token) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		//检查token 有效时间
		if !pkg.CheckTokenValidityTime(&tokenData) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is exceed", ""))
		}

		//刷新token有效时间
		if err = model.UserRefreshToken(db, id, tokenData); err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "db UserRefreshAppToken err", ""))
		}
		//if true {
		//	err := c.JSON("token???")
		//	if err != nil {
		//		return err
		//	}
		//	return nil
		//}
		//
		c.Locals(config.LOCAL_USERID_UINT, id)
		c.Locals(config.LOCAL_USERID_STRUCT, user)
		_ = c.Next()
		//c.JSON("231231231")
		return nil
	}
}

// AuthAppImg Protected protect routes
func AuthAppImg() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			id    uint
			token = c.Get(config.LOCAL_TOKEN)
			db    = database.DB
			err   error
		)

		// 打印请求地址
		log.Info("Request URL: ", c.Path())
		if token == "" || len(token) < 10 {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is null", ""))
		}

		id, tokenData, err := model.UserSelectAddressByToken(db, token)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		user := model.User{}
		user.ID = id
		err = user.GetById(db)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		if !pkg.CheckSpecialCharacters(&token) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		//检查token 有效时间
		if !pkg.CheckTokenValidityTime(&tokenData) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is exceed", ""))
		}

		//刷新token有效时间
		if err = model.UserRefreshToken(db, id, tokenData); err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "db UserRefreshAppToken err", ""))
		}
		//if true {
		//	err := c.JSON("token???")
		//	if err != nil {
		//		return err
		//	}
		//	return nil
		//}
		//
		c.Locals(config.LOCAL_USERID_UINT, id)
		c.Locals(config.LOCAL_USERID_STRUCT, user)
		_ = c.Next()
		//c.JSON("231231231")
		return nil
	}
}

// AuthManagerApp Protected protect routes
func AuthManagerApp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//var token = c.Get("token")
		//fmt.Println(token)
		//c.Locals("guo", name)
		//c.Locals("guo2", "23123123123123123")
		var (
			managerId uint
			token     = c.Get(config.LOCAL_TOKEN)
			db        = database.DB
			err       error
		)

		//开关 测试打开token
		//if true {
		//	c.Locals(config.LOCAL_USERID_UINT, 1)
		//	c.Locals(config.LOCAL_USERID_INT64, 1)
		//	_ = c.Next()
		//	return nil
		//}

		// 打印请求地址
		log.Info("Request URL: ", c.Path())
		if len(string(c.Body())) < 10000 {
			log.Info("Request JSON: ", string(c.Body()))
		}
		if token == "" || len(token) < 10 {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is null", ""))
		}

		managerId, tokenData, err := model.ManagerSelectIdByToken(db, token)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		if !pkg.CheckSpecialCharacters(&token) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		//检查token 有效时间
		if !pkg.CheckTokenValidityTime(&tokenData) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is exceed", ""))
		}

		//刷新token有效时间
		if err = model.ManagerRefreshToken(db, managerId, tokenData); err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "db UserRefreshAppToken err", ""))
		}
		//if true {
		//	err := c.JSON("token???")
		//	if err != nil {
		//		return err
		//	}
		//	return nil
		//}
		//
		c.Locals(config.LOCAL_MANAGERID_UINT, managerId)
		c.Locals(config.LOCAL_MANAGERID_INT64, int64(managerId))
		_ = c.Next()
		//c.JSON("231231231")
		return nil
	}
}

// ApiPrint
func ApiPrint() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println(c.BaseURL())
		fmt.Println(string(c.Body()))
		_ = c.Next()
		//c.JSON("231231231")
		return nil
	}
}

type CheckAmount struct {
	Amount string `json:"amount"`
}

// AuthCheckAmountApp 校验amount拦截器
func AuthCheckAmountApp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params CheckAmount

		// 打印请求地址
		log.Info("Request URL: ", c.Path())
		log.Info("Request JSON: ", string(c.Body()))

		// 解析请求体到 `CheckAmount` 结构体
		if err := c.BodyParser(&params); err != nil {
			log.Error("Failed to parse request body: ", err)
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Invalid request body", "请求错误"))
		}

		// 检查 `amount` 是否为正数
		var bi big.Int
		_, ok := bi.SetString(params.Amount, 10)
		if !ok {
			// 转换失败，说明 amount 不是有效的大整数字符串
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Amount must be a valid positive integer", "金额必须是一个有效的正整数"))
		}
		if bi.Sign() < 0 {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Amount must be a valid positive integer", "金额必须是一个有效的正整数"))
		}

		// 调用下一个中间件或处理程序
		if err := c.Next(); err != nil {
			log.Error("Next handler error: ", err)
			return err
		}

		return nil
	}
}

// AuthCheckFlagApp 校验flag拦截器
func AuthCheckFlagApp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals(config.LOCAL_USERID_STRUCT).(model.User)
		if user.Flag != "0" {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "user funds frozen", "用户资金已被冻结"))
		}

		// 调用下一个中间件或处理程序
		if err := c.Next(); err != nil {
			log.Error("Next handler error: ", err)
			return err
		}
		return nil
	}
}
