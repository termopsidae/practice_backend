package api

import (
	"fmt"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"log"
	c "paractice/config"
	"testing"
)

func TestTwitter(t *testing.T) {
	// 创建 OAuth1 配置
	config := oauth1.NewConfig(c.Config("CONSUMER_KEY"), c.Config("CONSUMER_SECRET"))
	// 创建访问令牌
	token := oauth1.NewToken(c.Config("ACCESS_TOKEN"), c.Config("ACCESS_TOKEN_SECRET"))
	// 使用配置和令牌创建 HTTP 客户端
	httpClient := config.Client(oauth1.NoContext, token)

	// 发起请求
	resp, err := httpClient.Post("https://api.twitter.com/oauth/request_token", "application/x-www-form-urlencoded", nil)
	if err != nil {
		log.Fatalf("Failed to post request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println("Response:", string(body))
}
