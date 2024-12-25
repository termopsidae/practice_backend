package types

// SubmitTradeOrderReq 挂单
type SubmitTradeOrderReq struct {
	Amount    string `json:"amount"`     // 总数量
	TradeType string `json:"trade_type"` // 交易类型 0-卖单 1-买单
	Price     string `json:"price"`      // 单价
}
type SubmitTradeOrderResp struct {
}

// CancelTradeOrderReq 取消挂单
type CancelTradeOrderReq struct {
	TradeOrderId uint `json:"trade_order_id"`
}
type CancelTradeOrderResp struct {
}

// AcceptTradeOrderReq 接受挂单
type AcceptTradeOrderReq struct {
	TradeOrderId uint   `json:"trade_order_id"`
	Amount       string `json:"amount"`
}
type AcceptTradeOrderResp struct {
}

// SelectTradeOrderListReq 查询挂单
type SelectTradeOrderListReq struct {
	OwnerOrNot     string   `json:"owner_or_not"`    // "1" - 查本用户 "0" - 查全市场
	TimeDuration   []int64  `json:"time_duration"`   // 时间筛选
	AmountDuration []string `json:"amount_duration"` // 剩余数量筛选
	PriceDuration  []string `json:"price_duration"`  // 单价筛选
	TradeType      string   `json:"trade_type"`      // 交易类型 0-卖单 1-买单 不传就所有
	Flag           string   `json:"flag"`            // 用户查询自己的时候填"0"-已取消 "1"-进行中 "2"-已完成 查全市场填 "1"-进行中 record(0,2)
	Page           int      `json:"page"`            // 分页
}
type SelectTradeOrderListResp struct {
	TradeOrderList []TradeOrderInfo `json:"trade_order_list"` // 单价
	PageSize       int              `json:"page_size"`
}

type TradeOrderInfo struct {
	ID                    uint                    `json:"id"`
	SellerAddress         string                  `json:"seller_address"`         // 挂单用户地址
	TeamName              string                  `json:"team_name"`              // 挂单用户团队名
	Price                 string                  `json:"price"`                  // 单价
	Amount                string                  `json:"amount"`                 // 总数量
	LastAmount            string                  `json:"last_amount"`            // 剩余数量
	TradeType             string                  `json:"trade_type"`             // 交易类型 0-卖单 1-买单
	AssociatedTransaction []AssociatedTransaction `json:"associated_transaction"` // 关联交易 只在查本用户时返回
	CreatedAt             string                  `json:"created_at"`             // 创建时间
	Top                   int64                   `json:"top"`                    // 0-取消置顶 1-置顶
	Flag                  string                  `json:"flag"`                   // 0-已取消 1-挂单中 2-已完成
	UserCreatedAt         string                  `json:"user_created_at"`        // 用户创建时间
}
type AssociatedTransaction struct {
	BuyerAddress string `json:"buyer_address"` // 购买用户地址
	Amount       string `json:"amount"`        // 购买数量
	PayPrice     string `json:"pay_price"`     // 购买支出
	Date         string `json:"time"`          // 交易时间
}

type GetTradeCommissionResp struct {
	TradeCommission string `json:"trade_commission"` // 手续费率
	TradeDiscount   string `json:"trade_discount"`   // 折扣率
	AIGPBuyLimit    string `json:"aigp_buy_limit"`   // 最低aigp买单数量
	AIGPSellLimit   string `json:"aigp_sell_limit"`  // 最低aigp卖单数量
	AIGPAvailable   string `json:"aigp_available"`   // 可用aigp余额
	USDTAvailable   string `json:"usdt_available"`   // 可用usdt余额
	AIGPPrice       string `json:"aigp_price"`       // 当前币价
}
