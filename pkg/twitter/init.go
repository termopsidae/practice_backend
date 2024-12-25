package twitter

// Client 客户端
var Client TwitterClient

func init() {

	newClient, err := NewTwitterClient()
	if err != nil {
		panic(err)
	}

	Client = newClient
}
