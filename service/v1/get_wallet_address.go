package service

import (
	"errors"
	"net/http"

	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/request"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/response"
)

type GetWalletAddressRequest struct {
	MerchantId int64  `json:"merchantId" mapstructure:"merchantId"`
	UserId     string `json:"userId" mapstructure:"userId"`
	Sign       string `json:"sign" mapstructure:"sign"`
}

type GetWalletAddressResponse []WalletAddressItem

type WalletAddressItem struct {
	Address  string `json:"address"`
	Currency string `json:"currency"`
	Network  string `json:"network"`
}

const getWalletAddressURI = "/getWalletAddress"

func NewGetWalletAddressRequest() *GetWalletAddressRequest {
	return &GetWalletAddressRequest{}
}

func (c *Client) GetWalletAddress(req *GetWalletAddressRequest) (resp *GetWalletAddressResponse, err error) {
	if req == nil {
		return nil, errors.New("getWalletAddress request is nil")
	}

	baseReq := &request.BaseRequest{
		URL:     getWalletAddressURI,
		Method:  http.MethodGet,
		Header:  nil,
		Version: sdk.VersionV1,
		Param:   req,
	}

	baseResp := response.NewBaseResponse(&GetWalletAddressResponse{})
	resp = baseResp.Data.(*GetWalletAddressResponse)
	err = c.Send(baseReq, baseResp)
	return
}
