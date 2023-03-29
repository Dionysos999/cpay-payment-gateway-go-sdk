package sdk

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"sync"

	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/log"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/request"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/response"
)

var defaultUserAgent = fmt.Sprintf("CPaySDK_Go/%s (%s; %s) Golang/%s", Version, runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"))

var (
	httpClient *http.Client
	loadOnce   sync.Once
)

// Client is the base struct of service clients
type Client struct {
	Credential  *Credential
	Config      *Config
	ServiceName string
	Logger      log.Logger
}

type SignFunc func(*http.Request) error

func (c *Client) Init(serviceName string) *Client {
	c.Logger = log.New()
	c.ServiceName = serviceName
	return c
}

func (c *Client) WithCredential(cred *Credential) *Client {
	c.Credential = cred
	return c
}

func (c *Client) WithSecret(mid int64, secretKey string) *Client {
	c.Credential = NewCredentials(mid, secretKey)
	return c
}

func (c *Client) WithConfig(config *Config) *Client {
	c.Config = config
	c.Logger.SetLevel(config.LogLevel)
	return c
}

// Send send the request and return the response to the client.
// Parameter request accepts concrete request object which follow Request.
func (c *Client) Send(req request.Request, resp response.Response) error {
	method := req.GetMethod()
	builder := GetParameterBuilder(method, c.Logger)

	mapReq, err := convStruct2Map(req.GetParam())
	if err != nil {
		return err
	}

	sign := GenSign(mapReq, c.Credential.SecretKey)
	mapReq["sign"] = sign

	encodedUrl, err := builder.BuildURL(req.GetURL(), mapReq)
	if err != nil {
		return err
	}

	endPoint := c.Config.Endpoint
	if endPoint == "" {
		endPoint = fmt.Sprintf("%s/%s", defaultTarget, c.ServiceName)
	}
	reqUrl := fmt.Sprintf("%s://%s/%s/%s%s", c.Config.Scheme, endPoint, c.ServiceName, req.GetVersion(), encodedUrl)

	body, err := builder.BuildBody(mapReq)
	if err != nil {
		return err
	}

	rawResponse, err := c.doSend(method, reqUrl, body, req.GetHeaders())
	if err != nil {
		return err
	}

	return response.ParseFromHttpResponse(rawResponse, resp)
}

func (c *Client) doSend(method, url, data string, header map[string]string) (*http.Response, error) {
	loadOnce.Do(func() {
		c.Logger.Infof("load once http client")
		httpClient = &http.Client{Timeout: c.Config.Timeout}
	})

	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		c.Logger.Errorf("%s", err.Error())
		return nil, err
	}

	c.setHeader(req, header)

	return httpClient.Do(req)
}

func (c *Client) setHeader(req *http.Request, header map[string]string) {
	if req.Method != http.MethodGet && req.Method != http.MethodDelete && req.Method != http.MethodHead {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Set("User-Agent", defaultUserAgent)

	for k, v := range header {
		req.Header.Set(k, v)
	}

	for k, v := range req.Header {
		c.Logger.Infof("header key: %s, header value: %s", k, v)
	}
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
