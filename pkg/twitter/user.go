package twitter

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const MeEndpoint = "/2/users/me"

func (c *twitterClient) GetUserMe(accessToken string) (*UserMe, error) {
	// 设置查询参数
	params := url.Values{}
	params.Add("user.fields", "name,pinned_tweet_id,profile_image_url,created_at,description,location,url,verified,withheld,entities,protected")
	//params.Add("expansions", "pinned_tweet_id")
	//params.Add("tweet.fields", "id,text")

	// 请求参数
	body, err := c.Get(MeEndpoint, accessToken, nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var responseData map[string]interface{}
	// unmarshal the json into a string map
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Printf("JSON error: %s", err)
		fmt.Println("JSON error: ", err.Error())
		return nil, err
	}
	fmt.Println("=========", responseData)
	// 解析响应 JSON
	var user GetUserMeResp
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Printf("JSON error: %s", err)
		return nil, err
	}
	return &user.Data, nil
}
