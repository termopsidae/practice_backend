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
