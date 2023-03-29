package example

import (
	"testing"
	"time"

	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/log"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/service/v1"
)

func TestCreateCreditCardOrder(t *testing.T) {
	client, _ := service.NewClientWithSecret(20000289, "b3pcduqzoi5ydsw4xa8064ne8e9h6x9j")

	client.WithConfig(sdk.NewDefaultSandBoxEnv().
		WithTimeout(time.Second * 5).
		WithLogLevel(log.DebugLevel))

	req := &service.CreateCreditCardOrderRequest{
		MerchantId:      20000289,
		MerchantTradeNo: "SDK_Test006",
		UserId:          "User_Test001",
		Currency:        "USD",
		Amount:          "1.50",
		Products:        `[{"name":"pen","price":"0.50","num":"3","currency":"USD"}]`,
		Country:         "US",
		Email:           "Devron90009@gmail.com",
		Ip:              "127.0.0.1",
		CallBackURL:     "https://www.google.com",
		SuccessURL:      "https://test.success.com",
		FailURL:         "https://test.fail.com",
		CreateTime:      time.Now().UnixMilli(),
	}

	resp, err := client.CreateCreditCardOrder(req)
	if err != nil {
		t.Fatal("err: ", err)
		return
	}

	t.Logf("get response body: %+v\n", resp)
}

func TestCreateCreditCardOrderCacheCustomerInfo(t *testing.T) {
	client, _ := service.NewClientWithSecret(20000289, "b3pcduqzoi5ydsw4xa8064ne8e9h6x9j")

	client.WithConfig(sdk.NewDefaultSandBoxEnv().
		WithTimeout(time.Second * 5).
		WithLogLevel(log.DebugLevel))

	req := &service.CreateCreditCardOrderRequest{
		MerchantId:      20000289,
		MerchantTradeNo: "SDK_Test007",
		UserId:          "User_Test002",
		Currency:        "USD",
		Amount:          "1.50",
		Products:        `[{"name":"pen","price":"0.5","num":"3","currency":"USD"}]`,
		Country:         "US",
		Email:           "Devron90009@gmail.com",
		Ip:              "127.0.0.1",
		CallBackURL:     "https://www.google.com",
		SuccessURL:      "https://test.success.com",
		FailURL:         "https://test.fail.com",
		CreateTime:      time.Now().UnixMilli(),
	}

	cReq := req.WithCacheCustomerInfo(&service.CreditCardCustomerInfo{
		AddressLine: "address-191.",
		City:        "Nicosia",
		Mobile:      "+35712345678",
		Zip:         "20ww",
	})

	resp, err := client.CreateCreditCardOrder(cReq)
	if err != nil {
		t.Fatal("err: ", err)
		return
	}

	t.Logf("get response body: %+v\n", resp)
}
