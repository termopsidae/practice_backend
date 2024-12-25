package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm/logger"
	"paractice/database"
	"paractice/model/api"
	"paractice/routing"
	"time"
)

func main() {
	database.ConnectDB()
	database.DB.Logger = database.DB.Logger.LogMode(logger.Warn)
	fiberApp := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024, // 设置文件上传大小限制为 50MB
	})

	// 添加 CORS 中间件
	fiberApp.Use(cors.New())
	// 使用 Limiter 中间件来限制请求速率
	fiberApp.Use(limiter.New(limiter.Config{
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // 根据 IP 进行限制
		},
		Max:        20,              // 每个 IP 每秒钟最多 20 次请求
		Expiration: 1 * time.Second, // 时间窗口为 1 秒
	}))
	// 添加 CORS 中间件
	fiberApp.Use(func(c *fiber.Ctx) error {
		// 允许所有域名进行跨域请求
		c.Set("Access-Control-Allow-Origin", "*")
		// 允许 GET、POST、PUT、DELETE 和 OPTIONS 方法进行跨域请求
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许客户端发送的请求头
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, token")
		// 在响应中添加 CORS 头
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		} else {
			return c.Next()
		}
	})
	err := api.InitUserTree(database.DB)
	if err != nil {
		panic(err)
	}

	routing.Setup(fiberApp)
	err = fiberApp.Listen(":4833")
	if err != nil {
		panic(err)
	}

}

func InitTask() {
	var (
		c = cron.New(cron.WithSeconds())
		//db  = database.DB
		err error
	)
	_, err = c.AddFunc("0 0 1 * * ?", func() {
		//api.IncomeRunP(database.DB)
		//api.ContractCycle(database.DB)
	})

	if err != nil {
		panic(err)
		return
	}
	c.Start()
}
