package types

// RegisterAndLoginReq 登录注册参数
type RegisterAndLoginReq struct {
	WalletAddress string `json:"wallet_address"`
	Message       string `json:"message"`   //签名的数据
	Signature     string `json:"signature"` //签名的结果
}
type RegisterAndLoginResp struct {
	Token string `json:"token"`
}

// SelectUserInfoReq 展示用户信息
type SelectUserInfoReq struct {
}
type SelectUserInfoResp struct {
	UserInfo UserInfo `json:"user_info"`
}
type UserInfo struct {
	WalletAddress                string `json:"wallet_address"`                  // 地址
	RecommendAddress             string `json:"recommend_address"`               // 推荐人地址
	TeamName                     string `json:"team_name"`                       // 用户团队名
	ABalance                     string `json:"a_balance"`                       // A币余额
	BBalance                     string `json:"b_balance"`                       // B币余额
	USDTBalance                  string `json:"usdt_balance"`                    // 充值USDT余额
	ExUSDTBalance                string `json:"ex_usdt_balance"`                 // 闪兑USDT余额
	Level                        int64  `json:"level"`                           // 等级
	Performance                  string `json:"performance"`                     // 除自己外总业绩
	Consume                      string `json:"consume"`                         // usdt消费
	PerformanceTotal             string `json:"performance_total"`               // 总业绩
	PerformanceIncrease          string `json:"performance_increase"`            // 昨日新增总业绩
	SmallZonePerformance         string `json:"small_zone_performance"`          // 小区业绩
	SmallZonePerformanceIncrease string `json:"small_zone_performance_increase"` // 昨日新增小区业绩
	NodeId                       uint   `json:"node_id"`                         // 0 无
	Flag                         string `json:"flag"`                            // 启用标志(0-正常 -1-冻结)
	USDTFund                     string `json:"usdt_fund"`                       // usdt资产
	AFund                        string `json:"a_fund"`                          // aigp资产
	CanWithdrawB                 string `json:"can_withdraw_b"`                  // 是否可以提现aigt(0-否 1-是)
	UpdateAt                     string `json:"updated_at"`
}

// BindRecommendIdReq 绑定推荐人
type BindRecommendIdReq struct {
	Uuid string `json:"uuid"` //推荐人地址
}
type BindRecommendIdResp struct {
}

// BindRecommenderReq 资产增值
type BindRecommenderReq struct {
	RecommenderAddress string `json:"recommender_address"`
}
type BindRecommenderResp struct {
}

type SengMsg struct {
	Msg   string `json:"msg"`
	Phone string `json:"phone"`
	Area  string `json:"area"`
}

// 查询所有商品列表接口
type SelectAllGoodsReq struct {
}
type SelectAllGoodsResp struct {
	Goods []GoodInfo `json:"goods"`
}
type GoodInfo struct {
	GoodId      uint    `json:"good_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	LastAmount  int64   `json:"last_amount"`
	Flag        string  `json:"flag"` // 启用标志(0-停用 1-可购买) `json:"goods"`
}
