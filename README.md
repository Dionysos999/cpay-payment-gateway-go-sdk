# CPay Payment Gateway Go SDK

`CPay Payment Gateway Go SDK` is written in Go.

# Quickstart

## Prerequisites
- Go v1.17 or heigher
- MerchantID and SecurityKey (from cpay team)

## Install
`go get github.com/cpayfinance/cpay-payment-gateway-go-sdk`

## Usage
The `example` package is a good place to start.

The minimal you'll need to have is:

``` go
// 1. create client by MerchantID and SecurityKey
client, _ := service.NewClientWithSecret(20000289, "b3pcduqzoi5ydsw4xa8064ne8e9h6x9j")

// 2. coufig your client

// a. use custom config
//	client.WithConfig(&sdk.Config{
//    Scheme:   sdk.SchemeHTTP,
//    Endpoint: `8.142.157.45:9075`,
//    Timeout:  time.Second * 10,
//    LogLevel: log.DebugLevel,
// })

// b. use default config
client.WithConfig(sdk.NewDefaultSandBoxEnv().
    WithTimeout(time.Second * 5).
    WithLogLevel(log.DebugLevel))

// 3. call the APIs
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

```
