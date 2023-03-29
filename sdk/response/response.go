package response

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/errors"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	TraceId string      `json:"traceid"`
}

var _ Response = &BaseResponse{}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) error {
	if err := json.Unmarshal(body, r); err != nil {
		return err
	}
	if r.Code > 0 {
		return errors.NewCPaySDKError(r.Code, r.Message, r.TraceId)
	}

	return nil
}

func NewBaseResponse(data interface{}) *BaseResponse {
	return &BaseResponse{
		Data: data,
	}
}

func ParseFromHttpResponse(rawResponse *http.Response, response Response) error {
	defer rawResponse.Body.Close()
	body, err := io.ReadAll(rawResponse.Body)
	if err != nil {
		return err
	}
	if rawResponse.StatusCode != 200 {
		return fmt.Errorf("request fail with status: %s, with body: %s", rawResponse.Status, body)
	}

	if err := response.ParseErrorFromHTTPResponse(body); err != nil {
		return err
	}
	return json.Unmarshal(body, &response)
}
