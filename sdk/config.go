package sdk

import (
	"time"
)

type Config struct {
	Scheme  string
	Timeout time.Duration
}

// NewConfig returns a pointer of Config
func NewConfig() *Config {
	return &Config{
		Scheme:  "HTTP",
		Timeout: 30 * time.Second,
	}
}

// WithScheme 协议
func (c *Config) WithScheme(scheme string) *Config {
	c.Scheme = scheme
	return c
}

// WithTimeout 超时时间
func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}
