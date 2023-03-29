package service

import (
	"errors"
	"net/http"

	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/request"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/response"
)

type CreateCryptoOrderRequest struct {
	MerchantId      int64  `json:"merchantId" mapstructure:"merchantId"`
	MerchantTradeNo string `json:"merchantTradeNo" mapstructure:"merchantTradeNo"`
	UserId          string `json:"userId" mapstructure:"userId"`
	CryptoCurrency  string `json:"cryptoCurrency" mapstructure:"cryptoCurrency"`
	Amount          string `json:"amount" mapstructure:"amount"`
	CreateTime      int64  `json:"createTime" mapstructure:"createTime"`
	CallBackURL     string `json:"callBackURL" mapstructure:"callBackURL"`
	ReturnURL       string `json:"returnURL" mapstructure:"returnURL"`
	SuccessURL      string `json:"successURL" mapstructure:"successURL"`
	FailURL         string `json:"failURL" mapstructure:"failURL"`
	Sign            string `json:"sign" mapstructure:"sign"`
}

type CreateCryptoOrderResponse struct {
	OrderId         string `json:"orderId"`
	MerchantId      string `json:"merchantId"`
	CryptoCurrency  string `json:"cryptoCurrency"`
	OrderAmount     string `json:"orderAmount"`
	ReceivedAmount  string `json:"receivedAmount"`
	MerchantTradeNo string `json:"merchantTradeNo"`
	CashierURL      string `json:"cashierURL"`
	ReturnURL       string `json:"returnURL"`
	SuccessURL      string `json:"successURL"`
	FailURL         string `json:"failURL"`
	Remark          string `json:"remark"`
	ExtInfo         string `json:"extInfo"`
}

const createOrderURI = "/createOrder"

func NewCreateOrderRequest() *CreateCryptoOrderRequest {
	return &CreateCryptoOrderRequest{}
}

func (c *Client) CreateCryptoOrder(req *CreateCryptoOrderRequest) (resp *CreateCryptoOrderResponse, err error) {
	if req == nil {
		return nil, errors.New("createCryptoOrder request is nil")
	}

	baseReq := &request.BaseRequest{
		URL:     createOrderURI,
		Method:  http.MethodPost,
		Header:  nil,
		Version: sdk.VersionV1,
		Param:   req,
	}

	baseResp := response.NewBaseResponse(&CreateCryptoOrderResponse{})
	resp = baseResp.Data.(*CreateCryptoOrderResponse)
	err = c.Send(baseReq, baseResp)
	return
}
