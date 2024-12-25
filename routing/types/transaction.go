package types

// ManagerUserTransactionInfoListReq  用户资产明细
type ManagerUserTransactionInfoListReq struct {
	WalletAddress string  `json:"wallet_address"` // 地址
	TimeDuration  []int64 `json:"time_duration"`  // 时间筛选
	ChangeType    string  `json:"change_type"`
	// 1-USDT充值 2-USDT提现 3-购买节点 4-AIGP提现 5-AIGT提现
	// 11-AIGP合成矿机消耗 12-USDT合成矿机消耗 13-USDT购买矿机消耗 14-AIGP转出 15-AIGP转入 16-AIGP兑换消耗 17-USDT兑换所得
	// 21-AIGT空投 22-AIGP挖矿产出 23-AIGP动态奖励 24-AIGP团队奖励 25-AIGP平级奖励 26-AIGP全球分红 27-AIGT每日销毁 28-AIGP手续费分红
	// 31-管理员增加用户余额 32-管理员扣减用户余额
	// USDT类变动 41-AIGP买单(-) 42-AIGP买单取消(+) 43-今日接受AIGP卖单支出(-) 44-接受AIGP买单收入(+) 45-AIGP卖单收入(+)
	// AIGP类变动 51-AIGP卖单(-) 52-AIGP卖单取消(+) 53-接受AIGP买单支出(-) 54-接受AIGP卖单收入(+) 55-AIGP买单收入(+)
	Flag string `json:"flag"` // 提现审核状态 0-审核中 2-审核通过 -1-审核驳回
	Page int    `json:"page"` // 页码从0开始
}

type ManagerUserTransactionInfoListResp struct {
	TransactionInfoList []TransactionInfo `json:"list"`
	PageSize            int               `json:"page_size"`
}

type UserTransactionInfoListReq struct {
	TimeDuration []int64 `json:"time_duration"` // 时间筛选
	ChangeType   string  `json:"change_type"`   // 1-USDT充值 withdraw-提现记录 transfer-转账记录 17-兑换记录 mining-资产明细 trade-市场交易记录
	Page         int     `json:"page"`          // 页码从0开始
}

type UserTransactionInfoListResp struct {
	TransactionInfoList []TransactionInfo `json:"list"`
	TotalPlus           string            `json:"total_plus"`    // 总增加
	TotalMinus          string            `json:"total_minus"`   // 总减少
	TotalUMinus         string            `json:"total_u_minus"` // u总减少
	TotalAMinus         string            `json:"total_a_minus"` // a总减少
	TotalBMinus         string            `json:"total_b_minus"` // b总减少
	PageSize            int               `json:"page_size"`     // 页长
}

type TransactionInfo struct {
	ID            uint   `json:"id"`             // 交易id
	WalletAddress string `json:"wallet_address"` // 地址
	Hash          string `json:"hash"`           // 冗余字段 14/15时为转出/转入地址 17时为消耗AIGP数量 22时为矿机等级
	Amount        string `json:"amount"`         // 数量
	ChangeType    string `json:"change_type"`
	// 1-USDT充值 2-USDT提现 3-购买节点 4-AIGP提现 5-AIGT提现
	// 11-AIGP合成矿机消耗 12-USDT合成矿机消耗 13-USDT购买矿机消耗 14-AIGP转出 15-AIGP转入 16-AIGP兑换消耗 17-USDT兑换所得
	// 21-AIGT空投 22-AIGP挖矿产出 23-AIGP动态奖励 24-AIGP团队奖励 25-AIGP平级奖励 26-AIGP全球分红 27-AIGT每日销毁 28-AIGP手续费分红
	// 31-管理员增加用户余额 32-管理员扣减用户余额
	// USDT类变动 41-AIGP买单(-) 42-AIGP买单取消(+) 43-今日接受AIGP卖单支出(-) 44-接受AIGP买单收入(+) 45-AIGP卖单收入(+)
	// AIGP类变动 51-AIGP卖单(-) 52-AIGP卖单取消(+) 53-接受AIGP买单支出(-) 54-接受AIGP卖单收入(+) 55-AIGP买单收入(+)
	Date string `json:"date"`
	Flag string `json:"flag"` // 提现审核状态 0-审核中 2-审核通过 -1-审核驳回
}
