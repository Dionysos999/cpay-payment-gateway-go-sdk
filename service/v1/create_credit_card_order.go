package service

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/request"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/response"
)

type CreateCreditCardOrderRequest struct {
	MerchantId        int64  `json:"merchantId" mapstructure:"merchantId"`
	MerchantTradeNo   string `json:"merchantTradeNo" mapstructure:"merchantTradeNo"`
	UserId            string `json:"userId" mapstructure:"userId"`
	Currency          string `json:"currency" mapstructure:"currency"`
	Amount            string `json:"amount" mapstructure:"amount"`
	Products          string `json:"products" mapstructure:"products"`
	Country           string `json:"country" mapstructure:"country"`
	Email             string `json:"email" mapstructure:"email"`
	Ip                string `json:"ip" mapstructure:"ip"`
	CallBackURL       string `json:"callBackURL" mapstructure:"callBackURL"`
	SuccessURL        string `json:"successURL" mapstructure:"successURL"`
	FailURL           string `json:"failURL" mapstructure:"failURL"`
	CreateTime        int64  `json:"createTime" mapstructure:"createTime"`
	CustomerCacheInfo string `json:"customerCacheInfo" mapstructure:"customerCacheInfo"` // 缓存用户基本信息,信用卡下单页面回显
	Sign              string `json:"sign" mapstructure:"sign"`
}

type CreditCardCustomerInfo struct {
	AddressLine string `json:"addressLine"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Zip         string `json:"zip"`
}

type CreateCreditCardOrderResponse struct {
	MerchantId      int64  `json:"merchantId"`
	MerchantTradeNo string `json:"merchantTradeNo"`
	OrderId         string `json:"orderId"`
	OrderStatus     string `json:"orderStatus"`
	OrderCurrency   string `json:"orderCurrency"`
	OrderAmount     string `json:"orderAmount"`
	ReceivedAmount  string `json:"receivedAmount"`
	PledgeAmount    string `json:"pledgeAmount"`
	Fee             string `json:"fee"`
	CashierURL      string `json:"cashierURL"`
	SuccessURL      string `json:"successURL"`
	FailURL         string `json:"failURL"`
}

const createOrderByCreditCardURI = "/createOrderByCreditCard"

func NewCreateCreditCardOrderRequest() *CreateCreditCardOrderRequest {
	return &CreateCreditCardOrderRequest{}
}

func (r *CreateCreditCardOrderRequest) WithCacheCustomerInfo(c *CreditCardCustomerInfo) *CreateCreditCardOrderRequest {
	marshal, err := json.Marshal(c)
	if err != nil {
		return r
	}
	r.CustomerCacheInfo = string(marshal)
	return r
}

func (c *Client) CreateCreditCardOrder(req *CreateCreditCardOrderRequest) (resp *CreateCreditCardOrderResponse, err error) {
	if req == nil {
		return nil, errors.New("createCreditCardOrder request is nil")
	}

	baseReq := &request.BaseRequest{
		URL:     createOrderByCreditCardURI,
		Method:  http.MethodPost,
		Header:  nil,
		Version: sdk.VersionV1,
		Param:   req,
	}

	baseResp := response.NewBaseResponse(&CreateCreditCardOrderResponse{})
	resp = baseResp.Data.(*CreateCreditCardOrderResponse)
	err = c.Send(baseReq, baseResp)
	return
}
