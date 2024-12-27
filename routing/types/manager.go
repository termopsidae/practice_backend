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
