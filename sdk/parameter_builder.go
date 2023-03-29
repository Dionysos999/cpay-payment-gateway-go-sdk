package sdk

import (
	"fmt"
	"net/http"
	urllib "net/url"
	"reflect"

	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/log"
	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/request"
)

var baseRequestFields []string

func init() {
	req := request.BaseRequest{}
	reqType := reflect.TypeOf(req)
	for i := 0; i < reqType.NumField(); i++ {
		baseRequestFields = append(baseRequestFields, reqType.Field(i).Name)
	}
}

type ParameterBuilder interface {
	BuildURL(url string, paramMap map[string]interface{}) (string, error)
	BuildBody(paramMap map[string]interface{}) (string, error)
}

func GetParameterBuilder(method string, logger log.Logger) ParameterBuilder {
	if method == http.MethodGet || method == http.MethodDelete || method == http.MethodHead {
		return &WithoutBodyBuilder{logger}
	} else {
		return &WithURLBodyBuilder{logger}
	}
}

// WithoutBodyBuilder supports GET/DELETE methods.
// It only builds path and query parameters.
type WithoutBodyBuilder struct {
	Logger log.Logger
}

func (b WithoutBodyBuilder) BuildURL(url string, paramMap map[string]interface{}) (res string, err error) {
	resultUrl := url
	values := urllib.Values{}
	if len(paramMap) > 0 {
		for k, v := range paramMap {
			values.Set(k, fmt.Sprintf("%v", v))
		}
		resultUrl += "?" + values.Encode()
	}

	b.Logger.Infof("URL=%s", resultUrl)
	return resultUrl, nil
}

func (b WithoutBodyBuilder) BuildBody(paramMap map[string]interface{}) (string, error) {
	return "", nil
}

type WithURLBodyBuilder struct {
	Logger log.Logger
}

func (b *WithURLBodyBuilder) BuildURL(url string, paramMap map[string]interface{}) (string, error) {
	b.Logger.Infof("URL=%s", url)
	return url, nil
}

func (b *WithURLBodyBuilder) BuildBody(paramMap map[string]interface{}) (string, error) {
	var res string
	values := urllib.Values{}
	if len(paramMap) > 0 {
		for k, v := range paramMap {
			values.Set(k, fmt.Sprintf("%v", v))
		}
		//res = urllib.QueryEscape(values.Encode())
		res = values.Encode()
	}
	b.Logger.Infof("Body=%s", res)
	return res, nil
}
