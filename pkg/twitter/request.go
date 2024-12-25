package twitter

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// request 发出 HTTP 请求，并返回响应体
func (c *twitterClient) request(method, endpoint string, body io.Reader, token string, extraHeaders map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, c.BaseURL+endpoint, body)
	if err != nil {
		return nil, err
	}

	// 设置默认请求头
	req.Header.Set("Content-Type", "application/json")

	// 特殊处理
	if endpoint == OAuth2Token {
		// 添加令牌到请求头
		if token != "" {
			req.Header.Set("Authorization", "Basic "+token)
		}
	} else {
		// 添加令牌到请求头
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
	}
	// 添加额外的请求头
	for key, value := range extraHeaders {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	fmt.Println("=====", string(all))
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	return all, err
}

// Get 发出 HTTP GET 请求
func (c *twitterClient) Get(endpoint string, token string, extraHeaders map[string]string, params url.Values) ([]byte, error) {
	// 将 URL 解析为一个 url.URL 对象
	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	// 将查询参数添加到 URL 对象
	url.RawQuery = params.Encode()

	// 调用 request 方法
	return c.request("GET", url.String(), nil, token, extraHeaders)
}

// Post 发出 HTTP POST 请求
func (c *twitterClient) Post(endpoint string, token string, body []byte, extraHeaders map[string]string) ([]byte, error) {
	return c.request("POST", endpoint, bytes.NewBuffer(body), token, extraHeaders)
}

// Delete 发出 HTTP DELETE 请求
func (c *twitterClient) Delete(endpoint string, token string, extraHeaders map[string]string) ([]byte, error) {
	return c.request("DELETE", endpoint, nil, token, extraHeaders)
}
