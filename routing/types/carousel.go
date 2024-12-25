package types

type Carousel struct {
	ID        uint   `json:"id"`
	EnContent string `json:"en_content"` // 英语文字内容
	ZhContent string `json:"zh_content"` // 繁体中文文字内容
	KrContent string `json:"kr_content"` // 韩语文字内容
	ViContent string `json:"vi_content"` // 越南语文字内容
	InContent string `json:"in_content"` // 印尼语文字内容
	MAContent string `json:"ma_content"` // 马来语文字内容
	Sort      int    `json:"sort"`       // 排序字段 越大越靠前
}

type ManagerQueryCarouselReq struct {
}

type ManagerQueryCarouselResp struct {
	CarouselList []Carousel `json:"carousel_list"`
}

type ManagerInsertNewCarouselReq struct {
	EnContent string `json:"en_content"` // 英语文字内容
	ZhContent string `json:"zh_content"` // 繁体中文文字内容
	KrContent string `json:"kr_content"` // 韩语文字内容
	ViContent string `json:"vi_content"` // 越南语文字内容
	InContent string `json:"in_content"` // 印尼语文字内容
	MAContent string `json:"ma_content"` // 马来语文字内容
}

type ManagerInsertNewCarouselResp struct {
}

type ManagerUpdateCarouselReq struct {
	UpdateCarouselList []Carousel `json:"update_carousel_list"`
}

type ManagerUpdateCarouselResp struct {
}

type ManagerDeleteCarouselReq struct {
	ID uint `json:"id"`
}

type ManagerDeleteCarouselResp struct {
}

type UserQueryCarouselReq struct {
}

type UserQueryCarouselResp struct {
	CarouselList []Carousel `json:"carousel_list"`
}
