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
	api.Post("/selectCartInfo", intcpt.AuthApp(), app.SelectCartInfo)

	//updateGoodOrder 商品订单修改 TODO
	api.Post("/updateGoodOrder", intcpt.AuthApp(), app.UpdateGoodOrder)

}
func ManageSetUp(api fiber.Router) {
	// login 管理员登录
	api.Post("/login", intcpt.ApiPrint(), app.Login)

	//createNewGood 创建新商品 （管理员创建新商品 设置商品名 价格 数量 等）
	api.Post("/createNewGood", intcpt.AuthManagerApp(), app.CreateNewGood)

	//selectGoodList 管理员查询商品列表 （管理员根据条件查询 1 商品名 2 是否可购买，3 是否有剩余 根据时间排序）

	//updateGoodInfo 更新商品信息 （管理员更新现有商品信息 商品名 描述，剩余数量，价格，状态）
}
