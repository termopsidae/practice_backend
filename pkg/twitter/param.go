package twitter

import "time"

type GetAccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type GetUserMeResp struct {
	Data UserMe `json:"data"`
}
type UserMe struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Username        string    `json:"username"`
	PinnedTweetID   string    `json:"pinned_tweet_id,omitempty"`
	ProfileImageURL string    `json:"profile_image_url,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	Description     string    `json:"description,omitempty"`
	Location        string    `json:"location,omitempty"`
	URL             string    `json:"url,omitempty"`
	Verified        bool      `json:"verified,omitempty"`
	Withheld        string    `json:"withheld,omitempty"`
	Entities        string    `json:"entities,omitempty"`
	Protected       bool      `json:"protected,omitempty"`
}

// Tweet represents a single Twitter post
type Tweet struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Language  string    `json:"lang"`
	Source    string    `json:"source"`
	Geo       TweetGeo  `json:"geo"`
}

// TweetGeo is a tweet location as a geometry (point in particular)
type TweetGeo struct {
	Coordinates Coordinate `json:"coordinates"`
}

// Coordinate is a coordinate object representing a point
type Coordinate struct {
	Type        string `json:"type"`
	Coordinates []int  `json:"coordinates"`
	PlaceID     string `json:"place_id"`
}

// RetweetedByQueryParams 封装 /2/tweets/:id/retweeted_by 端点的查询参数
type RetweetedByQueryParams struct {
	UserFields      []string // 用户字段列表
	Expansions      string   // 扩展字段
	TweetFields     []string // 推文字段列表
	MaxResults      int      // 最大结果数
	PaginationToken string   // 分页令牌
}

// FollowParams 结构用于存储关注推特账号所需的参数
type FollowParams struct {
	Username string `json:"username"`
}

// FollowResponse 结构用于存储关注推特账号的响应
type FollowResponse struct {
	Data struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		ID        string `json:"id"`
		Followers int    `json:"followers_count"`
	} `json:"data"`
}
