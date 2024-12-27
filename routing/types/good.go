package types

// PurchaseGoodReq 商品加入购物车
type PurchaseGoodReq struct {
	GoodId uint  `json:"good_id"`
	Amount int64 `json:"amount"`
}
type PurchaseGoodResp struct {
}

// SelectCartInfoReq 查询信息购物车
type SelectCartInfoReq struct {
}
type SelectCartInfoResp struct {
	CartId        uint            `json:"cart_id"`
	TotalPrice    float64         `json:"total_price"`
	GoodOrderList []GoodOrderInfo `json:"good_order_list"`
}

type GoodOrderInfo struct {
	GoodOrderId uint     `json:"good_order_id"`
	Good        GoodInfo `json:"good"`
	Amount      int64    `json:"amount"`
	TotalPrice  float64  `json:"total_price"`
}

// UpdateGoodOrderReq 更改订单信息
type UpdateGoodOrderReq struct {
	GoodOrderId uint  `json:"good_order_id"`
	Amount      int64 `json:"amount"` //更改后剩余数量 取消填0
}
type UpdateGoodOrderResp struct {
}
