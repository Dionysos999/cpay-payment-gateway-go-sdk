package example

import (
	"testing"
	"time"

	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/log"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/service/v1"
)

func TestGetWalletAddress(t *testing.T) {
	client, _ := service.NewClientWithSecret(20000205, "03dexh7nl4ywwaw2uuerge8ofxm4y0od")

	client.WithConfig(sdk.NewDefaultSandBoxEnv().
		WithTimeout(time.Second * 5).
		WithLogLevel(log.DebugLevel))

	resp, err := client.GetWalletAddress(&service.GetWalletAddressRequest{
		MerchantId: 20000205,
		UserId:     "User_SDK_Test001",
	})
	if err != nil {
		t.Fatal("err: ", err)
		return
	}

	t.Logf("get response body: %+v\n", resp)
}
