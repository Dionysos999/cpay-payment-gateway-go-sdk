package service

import (
	"errors"
	"net/http"

	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/request"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/response"
)

type GetPaymentOrderRequest struct {
	MerchantId      int64  `json:"merchantId" mapstructure:"merchantId"`
	MerchantTradeNo string `json:"merchantTradeNo" mapstructure:"merchantTradeNo"`
	CpayOrderId     string `json:"cpayOrderId" mapstructure:"cpayOrderId"`
	Hash            string `json:"hash" mapstructure:"hash"`
	Sign            string `json:"sign" mapstructure:"sign"`
}

type GetPaymentOrderResponse struct {
	MerchantId      int64  `json:"merchantId"`
	OrderId         string `json:"orderId"`
	Status          string `json:"status"`
	MerchantTradeNo string `json:"merchantTradeNo"`
	MerchantUserId  string `json:"merchantUserId"`
	Hash            string `json:"hash"`
	ActualAmount    string `json:"actualAmount"`
	ReceivedAmount  string `json:"receivedAmount"`
	PledgeAmount    string `json:"pledgeAmount"`
	Fee             string `json:"fee"`
	Currency        string `json:"currency"`
	Network         string `json:"network"`
	CreateTime      int64  `json:"createTime"`
	Remark          string `json:"remark"`
	ExtInfo         string `json:"extInfo"`
}

const getPaymentOrderURI = "/getOrderDetail"

func NewGetPaymentOrderRequest() *GetPaymentOrderRequest {
	return &GetPaymentOrderRequest{}
}

func (c *Client) GetPaymentOrder(req *GetPaymentOrderRequest) (resp *GetPaymentOrderResponse, err error) {
	if req == nil {
		return nil, errors.New("getPaymentOrder request is nil")
	}

	baseReq := &request.BaseRequest{
		URL:     getPaymentOrderURI,
		Method:  http.MethodGet,
		Header:  nil,
		Version: sdk.VersionV1,
		Param:   req,
	}

	baseResp := response.NewBaseResponse(&GetPaymentOrderResponse{})
	resp = baseResp.Data.(*GetPaymentOrderResponse)
	err = c.Send(baseReq, baseResp)
	return
}
