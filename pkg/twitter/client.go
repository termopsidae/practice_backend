package twitter

import (
	"crypto/tls"
	"net"
	"net/http"
	"paractice/config"
	"time"
)

type twitterClient struct {
	Scopes        string
	ClientID      string
	ClientSecret  string
	CodeChallenge string
	HTTPClient    *http.Client
	BaseURL       string
}
type TwitterClient interface {
	GenerateAuthURL(redirectUrl string) string
	GetAccessToken(authorizationCode string, redirectUrl string) (*GetAccessTokenResp, error)
	GetUserMe(accessToken string) (*UserMe, error)
	FollowUser(screenName, accessToken string) error
	ReTweet(tweetID int64, accessToken string) error
}

func NewTwitterClient() (TwitterClient, error) {
	c := &twitterClient{
		Scopes:        config.Config("SCOPES"),
		ClientID:      config.Config("CLIENT_ID"),
		ClientSecret:  config.Config("CLIENT_SECRET"),
		CodeChallenge: config.Config("CODE_CHALLENGE"),
		BaseURL:       config.Config("BASE_URL"),
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 60 * time.Second,
				}).DialContext,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				MaxIdleConns:          0,
				MaxIdleConnsPerHost:   3000,
				MaxConnsPerHost:       3000,
				IdleConnTimeout:       60 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 3 * time.Second,
			},
		},
	}
	return c, nil
}
