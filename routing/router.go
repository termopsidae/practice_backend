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

	//selectAllGoods 查询所有商品接口
	api.Post("/selectAllGoods", app.SelectAllGoods)

	//purchaseGood 商品加入购物车 TODO
	api.Post("/purchaseGood", intcpt.AuthApp(), app.PurchaseGood)

	//selectCartInfo 查询购物车 TODO 仅查询待付款
	api.Post("/selectCartInfo", intcpt.AuthApp(), app.SelectAllGoods)

	//updateGoodOrder 商品订单修改 TODO
	api.Post("/updateGoodOrder", intcpt.AuthApp(), app.SelectAllGoods)

}
func ManageSetUp(api fiber.Router) {
	// login 管理员登录
	api.Post("/login", intcpt.ApiPrint(), app.Login)

}
