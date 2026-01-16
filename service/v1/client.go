package service

import "github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk"

const (
	defaultEndpoint = "https://api.cpay.ltd"
	serviceName     = "openapi"
)

type Client struct {
	sdk.Client
}

func NewClient(config *sdk.Config, credential *sdk.Credential) *Client {
	client := &Client{}
	if config == nil {
		config = sdk.NewConfig().WithEndpoint(defaultEndpoint)
	}

	client.Init(serviceName).WithCredential(credential).WithConfig(config)
	return client
}

func NewClientWithSecret(mid int64, secretKey string) *Client {
	client := &Client{}
	config := sdk.NewConfig().WithEndpoint(defaultEndpoint)
	client.Init(serviceName).WithSecret(mid, secretKey).WithConfig(config)
	return client
}
