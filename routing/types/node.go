package types

// PurchaseNodeReq 入金购买Node
type PurchaseNodeReq struct {
	NodeId uint   `json:"node_id"` //todo
	Hash   string `json:"hash"`
}
type PurchaseNodeResp struct {
}

// SelectNodeListReq 查询矿机列表
type SelectNodeListReq struct {
}
type SelectNodeListResp struct {
	NodeList   []NodeInfo `json:"node_list"`
	CanBuyNode string     `json:"can_buy_node"` // 是否可以前往购买节点 0-否 1是
}

type NodeInfo struct {
	NodeId   uint   `json:"node_id"`
	Name     string `json:"node_name"` // 节点名称
	Price    string `json:"price"`     // 价格 bi
	Percent  string `json:"percent"`   // 手续费分红比例
	Last     string `json:"last"`      // 剩余数量
	Level    int64  `json:"level"`     // 等级
	Max      string `json:"max"`       // 最大数量
	BBenefit string `json:"b_benefit"` // 空投AIGT数量 bi
	MiningId uint   `json:"mining_id"` // 赠送矿机id
}

// SelectNodeOrderListReq 查询节点列表
type SelectNodeOrderListReq struct {
	WalletAddress string `json:"wallet_address"` //地址
	NodeId        uint   `json:"node_id"`
	Flag          string `json:"flag"` // 0-考核通过 -1-考核未通过
	Page          int    `json:"page"` //分页
}
type SelectNodeOrderListResp struct {
	NodeList []NodeOrderInfo `json:"node_order_list"`
	Summary  string          `json:"summary"`
	PageSize int             `json:"page_size"`
}
type NodeOrderInfo struct {
	WalletAddress string `json:"wallet_address"` //地址
	Price         string `json:"price"`          // 售价
	Percent       string `json:"percent"`        // 分红比例
	CreatedAt     string `json:"created_at"`     // 创建日期
	Flag          string `json:"flag"`           // 0-考核通过 -1-考核未通过
}
