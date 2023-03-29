package errors

import "fmt"

type CPaySDKError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	TraceId string `json:"traceid"`
}

func (err *CPaySDKError) Error() string {
	return fmt.Sprintf("[CPaySDKError] code=%d, message=%s, traceid=%s", err.Code, err.Message, err.TraceId)
}

func NewCPaySDKError(code int, msg, traceId string) error {
	return &CPaySDKError{
		Code:    code,
		Message: msg,
		TraceId: traceId,
	}
}

func (err *CPaySDKError) GetCode() int {
	return err.Code
}

func (err *CPaySDKError) GetMessage() string {
	return err.Message
}

func (err *CPaySDKError) GetTraceID() string {
	return err.TraceId
}
