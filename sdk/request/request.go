package request

type Request interface {
	GetURL() string
	GetMethod() string
	GetVersion() string
	GetHeaders() map[string]string
	GetParam() interface{}
}

// BaseRequest is the base struct of service requests
type BaseRequest struct {
	URL     string
	Method  string
	Header  map[string]string
	Version string
	Param   interface{}
}

func (r BaseRequest) GetURL() string {
	return r.URL
}

func (r BaseRequest) GetMethod() string {
	return r.Method
}

func (r BaseRequest) GetVersion() string {
	return r.Version
}

func (r BaseRequest) GetHeaders() map[string]string {
	return r.Header
}

func (r BaseRequest) GetParam() interface{} {
	return r.Param
}
