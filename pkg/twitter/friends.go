package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func (c *twitterClient) FollowUser(screenName, accessToken string) error {
	config := oauth1.NewConfig(c.ClientID, c.ClientSecret)
	token := oauth1.NewToken(accessToken, "")
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	_, _, err := client.Friendships.Create(&twitter.FriendshipCreateParams{
		ScreenName: screenName,
	})
	return err
}
func (c *twitterClient) ReTweet(tweetID int64, accessToken string) error {
	config := oauth1.NewConfig(c.ClientID, c.ClientSecret)
	token := oauth1.NewToken(accessToken, "")
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	_, _, err := client.Statuses.Retweet(tweetID, &twitter.StatusRetweetParams{})
	return err
}
