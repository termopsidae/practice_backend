package types

// ShowCurrencyLogResp 币价趋势
type ShowCurrencyLogResp struct {
	Dates                []string `json:"values"`
	YesterdayAllAIGPCost string   `json:"yesterday_all_aigp_cost"` // 昨日全网aigp销毁数
}

// ShowCurrencyReq 币价
type ShowCurrencyReq struct {
	CoinType string `json:"coin_type"` // 币种 A-AIGP
}
type ShowCurrencyResp struct {
	Value string `json:"value"`
}

// TransferABalanceReq 转账
type TransferABalanceReq struct {
	ToAddress string `json:"to_address"`
	Amount    string `json:"amount"`
}
type TransferABalanceResp struct {
}

// ChangeAToUReq 换币
type ChangeAToUReq struct {
	Amount string `json:"amount"` // 金额
}

type ChangeAToUResp struct {
}

type AToUExceptedResp struct {
	Amount string `json:"amount"` // 实得金额
	Cost   string `json:"cost"`   // 手续费
}

type GetExchangeCommissionResp struct {
	ExchangeCommission string `json:"exchange_commission"` // 费率
}
type GetWithdrawCommissionResp struct {
	WithdrawCommission string `json:"withdraw_commission"` // 费率
	CanWithdrawA       string `json:"can_withdraw_a"`      // 是否可以提现aigp(0-否 1-是)
	CanWithdrawB       string `json:"can_withdraw_b"`      // 是否可以提现aigt(0-否 1-是)
	UWithdrawLimit     string `json:"u_withdraw_limit"`    // u提现限额
	AWithdrawLimit     string `json:"a_withdraw_limit"`    // a提现限额
	BWithdrawLimit     string `json:"b_withdraw_limit"`    // b提现限额
}
