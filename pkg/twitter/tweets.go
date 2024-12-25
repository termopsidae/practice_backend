package twitter

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	TweetRetweeted = "/2/tweets/%s/retweeted_by"
)

// GetTweetRetweetedBy 获取被指定推文 ID 转推的用户列表
func (c *twitterClient) GetTweetRetweetedBy(tweetID string, token string, queryParams RetweetedByQueryParams) ([]byte, error) {
	// 定义 endpoint
	endpoint := fmt.Sprintf(TweetRetweeted, tweetID)

	// 将 queryParams 转换为查询参数
	params := url.Values{}

	// 转换 UserFields 参数
	if len(queryParams.UserFields) > 0 {
		params.Set("user.fields", strings.Join(queryParams.UserFields, ","))
	}

	// 设置其他查询参数
	if queryParams.Expansions != "" {
		params.Set("expansions", queryParams.Expansions)
	}
	if len(queryParams.TweetFields) > 0 {
		params.Set("tweet.fields", strings.Join(queryParams.TweetFields, ","))
	}
	if queryParams.MaxResults > 0 {
		params.Set("max_results", fmt.Sprintf("%d", queryParams.MaxResults))
	}
	if queryParams.PaginationToken != "" {
		params.Set("pagination_token", queryParams.PaginationToken)
	}

	// 调用客户端的 Get 方法
	return c.Get(endpoint, token, nil, params)
}
