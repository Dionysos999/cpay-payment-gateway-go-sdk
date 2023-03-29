package example

import (
	"testing"
	"time"

	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/log"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/service/v1"
)

func TestCreateCryptoOrder(t *testing.T) {
	client, _ := service.NewClientWithSecret(20000289, "b3pcduqzoi5ydsw4xa8064ne8e9h6x9j")

	// a. custom config
	/*	client.WithConfig(&sdk.Config{
		Scheme:   sdk.SchemeHTTP,
		Endpoint: `8.142.157.45:9075`,
		Timeout:  time.Second * 10,
		LogLevel: log.DebugLevel,
	})*/

	// b. use default config
	client.WithConfig(sdk.NewDefaultSandBoxEnv().
		WithTimeout(time.Second * 5).
		WithLogLevel(log.DebugLevel))

	resp, err := client.CreateCryptoOrder(&service.CreateCryptoOrderRequest{
		MerchantId:      20000289,
		MerchantTradeNo: "SDK_Test004",
		UserId:          "User_Test002",
		CryptoCurrency:  "USDT",
		Amount:          "10.0",
		CreateTime:      time.Now().UnixMilli(),
		CallBackURL:     "https://www.google.com",
		SuccessURL:      "https://test.success.com",
		FailURL:         "https://test.fail.com",
	})
	if err != nil {
		t.Fatal("err: ", err)
		return
	}

	t.Logf("get response body: %+v\n", resp)
}
