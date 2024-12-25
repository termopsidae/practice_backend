package twitter

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const (
	OAuth2Token = "/2/oauth2/token"
)

func (c *twitterClient) GenerateAuthURL(redirectUrl string) string {
	authorizationURL := fmt.Sprintf(
		"https://twitter.com/i/oauth2/authorize?"+
			"scope=%s"+
			"&state=state"+
			"&response_type=code&client_id=%s"+
			"&code_challenge=%s"+
			"&code_challenge_method=plain&redirect_uri=%s",
		c.Scopes, c.ClientID, c.CodeChallenge, redirectUrl)

	return authorizationURL
}

func (c *twitterClient) GetAccessToken(authorizationCode string, redirectUrl string) (*GetAccessTokenResp, error) {
	// set the url and form-encoded data for the POST to the access token endpoint
	data := fmt.Sprintf(
		"grant_type=authorization_code&client_id=%s"+
			"&code_verifier=%s"+
			"&code=%s"+
			"&redirect_uri=%s",
		c.ClientID, c.CodeChallenge, authorizationCode, redirectUrl)
	fmt.Println("data", data)
	header := map[string]string{
		"content-type": "application/x-www-form-urlencoded",
	}
	// Encode username:password combination using base64 encoding as new bearer token to acquire access token
	newBearerToken := base64.StdEncoding.EncodeToString([]byte(c.ClientID + ":" + c.ClientSecret))
	body, err := c.Post(OAuth2Token, newBearerToken, []byte(data), header)
	if err != nil {
		fmt.Printf("HTTP error: %s", err)
		return nil, err
	}

	// process the response
	var responseData GetAccessTokenResp
	// unmarshal the json into a string map
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Printf("JSON error: %s", err)
		return nil, err
	}
	fmt.Println("=========", responseData)
	// retrieve the access token out of the map, and return to caller
	return &responseData, nil
}
