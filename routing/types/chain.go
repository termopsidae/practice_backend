package types

type RechargeConfirmReq struct {
	Hash string `json:"hash"` //交易hash
}
type RechargeConfirmResp struct {
}

// GetRechargeTransactionDataReq 充值
type GetRechargeTransactionDataReq struct {
	TokenName string `json:"token_name"` // USDT AIGP AIGT
	Num       string `json:"num"`        // 输入案例 "1.22" = 1.22 USDT
}

type WithdrawalReq struct {
	TokenName        string `json:"token_name"`        // USDT AIGP AIGT
	Num              string `json:"num"`               // 输入案例 "1.22" = 1.22 USDT
	MessageSignature string `json:"message_signature"` //消息签名
	Message          string `json:"message"`           //消息 “0x31321312321 提现 100 USDT”
}

type RechargeContractUSDCInfoResp struct {
	ToAddress string `json:"to_address"`
	USDCWei   string `json:"usdc_wei"`
	USDT      string `json:"usdt"`
}
