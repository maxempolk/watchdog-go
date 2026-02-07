package endpoint

import (
	"fmt"
	"strings"
	"time"
)

const (
	MethodHead = "HEAD"
	MethodGet  = "GET"
	MethodPost = "POST"

	DefaultTimeout = 3 * time.Second
	TrendSize      = 4
)

type EndpointConfig struct {
	URL                  string
	Interval             time.Duration
	Method               string
	Headers              map[string]string
	Body                 string
	ExpectedStatus       int
	ExpectedBodyContains string
	Timeout              time.Duration
	FollowRedirects      bool
}

type Endpoint struct {
	URL                  string
	Interval             time.Duration
	Method               string
	Headers              map[string]string
	Body                 string
	ExpectedStatus       int
	ExpectedBodyContains string
	Timeout              time.Duration
	FollowRedirects      bool

	Status    int
	Latency   time.Duration
	LastCheck time.Time
	Trend     []bool
}

func NewDefaultEndpointConfig(url string, interval time.Duration) EndpointConfig {
	return EndpointConfig{
		URL:             url,
		Interval:        interval,
		Method:          MethodHead,
		Headers:         map[string]string{},
		Timeout:         DefaultTimeout,
		FollowRedirects: true,
	}
}

func NormalizeEndpointConfig(cfg EndpointConfig) (EndpointConfig, error) {
	cfg.URL = strings.TrimSpace(cfg.URL)
	if cfg.URL == "" {
		return EndpointConfig{}, fmt.Errorf("url is required")
	}

	if cfg.Interval <= 0 {
		return EndpointConfig{}, fmt.Errorf("interval must be > 0")
	}

	cfg.Method = strings.ToUpper(strings.TrimSpace(cfg.Method))
	if cfg.Method == "" {
		cfg.Method = MethodHead
	}
	if cfg.Method != MethodHead && cfg.Method != MethodGet && cfg.Method != MethodPost {
		return EndpointConfig{}, fmt.Errorf("method must be one of: HEAD, GET, POST")
	}

	if cfg.Timeout == 0 {
		cfg.Timeout = DefaultTimeout
	}
	if cfg.Timeout < 0 {
		return EndpointConfig{}, fmt.Errorf("timeout must be >= 0")
	}

	if cfg.ExpectedStatus != 0 && (cfg.ExpectedStatus < 100 || cfg.ExpectedStatus > 599) {
		return EndpointConfig{}, fmt.Errorf("expected_status must be 0 or a valid HTTP status code")
	}

	cfg.ExpectedBodyContains = strings.TrimSpace(cfg.ExpectedBodyContains)

	cfg.Headers = copyHeaders(cfg.Headers)
	for k := range cfg.Headers {
		if strings.TrimSpace(k) == "" {
			return EndpointConfig{}, fmt.Errorf("headers must not contain empty keys")
		}
	}

	return cfg, nil
}

func NewEndpoint(cfg EndpointConfig) Endpoint {
	return Endpoint{
		URL:                  cfg.URL,
		Interval:             cfg.Interval,
		Method:               cfg.Method,
		Headers:              copyHeaders(cfg.Headers),
		Body:                 cfg.Body,
		ExpectedStatus:       cfg.ExpectedStatus,
		ExpectedBodyContains: cfg.ExpectedBodyContains,
		Timeout:              cfg.Timeout,
		FollowRedirects:      cfg.FollowRedirects,
		Trend:                make([]bool, TrendSize),
	}
}

func (e Endpoint) ToConfig() EndpointConfig {
	return EndpointConfig{
		URL:                  e.URL,
		Interval:             e.Interval,
		Method:               e.Method,
		Headers:              copyHeaders(e.Headers),
		Body:                 e.Body,
		ExpectedStatus:       e.ExpectedStatus,
		ExpectedBodyContains: e.ExpectedBodyContains,
		Timeout:              e.Timeout,
		FollowRedirects:      e.FollowRedirects,
	}
}

func copyHeaders(src map[string]string) map[string]string {
	if len(src) == 0 {
		return map[string]string{}
	}

	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func (e Endpoint) IsHealthy() bool {
	return e.Status >= 200 && e.Status < 400
}

type HealthCheckResult struct {
	URL     string
	Status  int
	Latency time.Duration
}

type HealthChecker interface {
	Check(endpoint Endpoint) (HealthCheckResult, error)
}
