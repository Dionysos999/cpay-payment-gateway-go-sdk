package sdk

import (
	"time"

	"github.com/cpayfinance/cpay-payment-gateway-go-sdk/sdk/log"
)

type Config struct {
	Scheme   string
	Endpoint string
	Timeout  time.Duration
	LogLevel log.Level
}

var defaultTarget = "api.cpay.ltd"

// NewConfig returns a pointer of Config
// scheme only accepts http or https
func NewConfig() *Config {
	return &Config{
		Scheme:   SchemeHTTP,
		Timeout:  15 * time.Second,
		LogLevel: log.WarnLevel,
	}
}

func NewDefaultSandBoxEnv() *Config {
	return &Config{
		Scheme:   SchemeHTTP,
		Endpoint: "8.142.157.45:9075",
	}
}

func NewDefaultLiveEnv() *Config {
	return &Config{
		Scheme:   SchemeHTTPS,
		Endpoint: defaultTarget,
	}
}

func (c *Config) WithScheme(scheme string) *Config {
	c.Scheme = scheme
	return c
}

func (c *Config) WithEndpoint(endpoint string) *Config {
	c.Endpoint = endpoint
	return c
}

func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}

func (c *Config) WithLogLevel(level log.Level) *Config {
	c.LogLevel = level
	return c
}
