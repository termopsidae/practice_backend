package types

type MiningInfo struct {
	MiningId      uint   `json:"mining_id"`
	Name          string `json:"nft_name"`      // 矿机名称
	CoinType      string `json:"coin_type"`     // 产币类型
	Price         string `json:"price"`         // 价格 bi
	Magnification string `json:"magnification"` // 倍率
	AllIncome     string `json:"all_income"`    // 总收益 bi
	BBenefit      string `json:"b_benefit"`     // 空投AIGT数量 bi
	Profit        string `json:"profit"`        // 释放率
	Flag          string `json:"flag"`          // 启用标志(0-停用 1-可合成 2-可购买 3-可合成可购买)
}

// PurchaseNftReq 入金购买nft
type PurchaseNftReq struct {
	MiningId uint   `json:"mining_id"`
	Hash     string `json:"hash"`
}
type PurchaseNftResp struct {
}

// SelectMiningListReq 查询矿机列表
type SelectMiningListReq struct {
}
type SelectMiningListResp struct {
	MiningList  []MiningInfo `json:"mining_list"`
	USDTPercent string       `json:"usdt_percent"`
}

// SelectMiningOrderListReq 查询矿机订单列表
type SelectMiningOrderListReq struct {
	WalletAddress string `json:"wallet_address"` // 地址
	Level         uint   `json:"level"`          // 矿机等级
	Flag          string `json:"flag"`           // 1-进行中 0-已终止
	Page          int    `json:"page"`           // 分页
}
type SelectMiningOrderListResp struct {
	MiningOrderList []MiningOrderInfo `json:"mining_order_list"`
	PageSize        int               `json:"page_size"`
}
type MiningOrderInfo struct {
	MiningOrderId uint   `json:"mining_order_id"`
	WalletAddress string `json:"wallet_address"` // 地址
	Price         string `json:"price"`          // 价格
	AllIncome     string `json:"all_income"`     // 总收益
	BBenefit      string `json:"b_benefit"`      // 空投AIGT
	Profit        string `json:"profit"`         // 每日释放
	ReleaseNum    string `json:"release_num"`    // 未释放
	Flag          string `json:"flag"`           // 状态 1-进行中 0-已终止
	Level         uint   `json:"level"`          // 等级
	CreatedAt     string `json:"created_at"`     // 创建时间
}

// MiningSynthesisReq 合成矿机
type MiningSynthesisReq struct {
	MiningId uint `json:"mining_id"`
	//AIGPAll  string `json:"aigp_all"` // 0-正常合成 1-全部使用a币
}

type MiningSynthesisResp struct {
}

// MiningSynthesisExceptedReq 合成矿机预期
type MiningSynthesisExceptedReq struct {
	MiningId uint `json:"mining_id"`
}

type MiningSynthesisExceptedResp struct {
	USDTExcepted string `json:"usdt_excepted"`
	AIGPExcepted string `json:"aigp_excepted"`
	//CanUseAllA      string `json:"can_use_all_a"` // 0-否 1-可全部使用a币
	//AIGPAllExcepted string `json:"aigp_all_excepted"`
}

// PurchaseMiningReq 购买矿机
type PurchaseMiningReq struct {
	MiningId uint `json:"mining_id"`
}
type PurchaseMiningResp struct {
}

type UserMiningList struct {
	MiningList []MiningOrderInfo `json:"mining_order_list"`
}
