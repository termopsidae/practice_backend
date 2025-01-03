package types

// LoginReq 登录注册参数
type LoginReq struct {
	ManagerName string `json:"manager_name"` //管理员用户名
	Password    string `json:"password"`     //密码
}
type LoginResp struct {
	Token string `json:"token"` // token
	Class string `json:"class"` // 管理员等级
}

// 管理员创建商品
type CreatNewGoodReq struct {
	GoodName    string  `json:"good_name"`   //商品名
	Price       float64 `json:"price"`       //商品价格
	LastAmount  int64   `json:"last_amount"` //商品数量
	Description string  `json:"description"` //商品描述
}
type CreatNewGoodResp struct {
}

// 管理员根据条件查询商品列表
type SelectGoodListReq struct {
	SelectCondition int64  `json:"select_condition"` //以什么样的条件查询 1 商品名 2 是否可购买，3 是否有剩余
	SelectValue     string `json:"select_value"`     //查询的值

}
type SelectGoodListResp struct {
	GoodsList []TypeGood `json:"goods_list"`
}
type TypeGood struct {
	GoodName    string  `json:"good_name"`   //商品名
	Price       float64 `json:"price"`       //商品价格
	LastAmount  int64   `json:"last_amount"` //商品数量
	Description string  `json:"description"` //商品描述
	Flag        string  `json:"flag"`        //商品售卖状态  1 销售中. 2暂停销售
}
