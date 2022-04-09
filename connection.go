package violetlog

import "gopkg.in/h2non/gentleman.v2"

type VioletLog struct {
	client *gentleman.Client
	url    string
}

func New(url, token string) VioletLog {
	client := gentleman.New()
	client.BaseURL(url)
	client.AddHeader("Authorization", token)

	return VioletLog{
		client: client,
	}
}
