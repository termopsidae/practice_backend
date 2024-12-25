package routing

import (
	"github.com/gofiber/fiber/v2"
	"paractice/routing/app"
	intcpt "paractice/routing/intercept"
)

func Setup(f *fiber.App) {
	appApi := f.Group("/app")
	managerApi := f.Group("/manage")
	AppSetUp(appApi)
	ManageSetUp(managerApi)

}
func AppSetUp(api fiber.Router) {
	// registerAndLogin 登录注册
	api.Post("/registerAndLogin", app.RegisterAndLogin)

}
func ManageSetUp(api fiber.Router) {
	// login 管理员登录
	api.Post("/login", intcpt.ApiPrint(), app.Login)

}
