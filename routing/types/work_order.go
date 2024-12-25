package types

// UploadImgReq 上传图片参数
type UploadImgReq struct {
	Hash string `json:"hash"` //工单编号
	Data string `json:"data"`
}
type UploadImgResp struct {
	URLs []string `json:"urls"`
}

// UploadWorkOrderReq 上传工单参数
type UploadWorkOrderReq struct {
	Hash                  string `json:"hash"`                     // 工单编号
	ProblemImageUrlAssets string `json:"problem_image_url_assets"` // 问题图片url链接
	ProblemDescription    string `json:"problem_description"`      // 问题文字描述
	TelegramNumber        string `json:"telegram_number"`          // tg号(选填)
}
type UploadWorkOrderResp struct {
}

// UploadWorkOrderReplyReq 回复工单参数
type UploadWorkOrderReplyReq struct {
	Hash                string `json:"hash"`                   // 工单编号
	ReplyImageUrlAssets string `json:"reply_image_url_assets"` // 问题图片url链接
	ReplyDescription    string `json:"reply_description"`      // 问题文字描述
}
type UploadWorkOrderReplyResp struct {
}

// SelectWorkOrderListReq  查询工单
type SelectWorkOrderListReq struct {
	Hash          string `json:"hash"`           // 工单编号
	WalletAddress string `json:"wallet_address"` // 地址
	Flag          string `json:"flag"`
	Page          int    `json:"page"` // 页码从0开始
}
type SelectWorkOrderListResp struct {
	WorkOrderList []WorkOrderInfo `json:"work_order_list"`
	PageSize      int             `json:"page_size"`
}
type WorkOrderInfo struct {
	Hash                  string `json:"hash"`                     // 工单编号
	WalletAddress         string `json:"wallet_address"`           // 地址
	ProblemImageUrlAssets string `json:"problem_image_url_assets"` // 问题图片url链接
	ProblemDescription    string `json:"problem_description"`      // 问题文字描述
	ReplyImageUrlAssets   string `json:"reply_image_url_assets"`   // 回复图片url链接
	ReplyDescription      string `json:"reply_description"`        // 回复文字描述
	TelegramNumber        string `json:"telegram_number"`          // tg号(选填)
	Flag                  string `json:"flag"`                     // 0 正在处理 1 处理完成
	CreatedAt             string `json:"created_at"`               // 创建时间
}

// UserQueryWorkOrderListReq  用户查询工单
type UserQueryWorkOrderListReq struct {
}

type UserQueryWorkOrderListResp struct {
	WorkOrderList []WorkOrderInfo `json:"work_order_list"`
	PageSize      int             `json:"page_size"`
}
