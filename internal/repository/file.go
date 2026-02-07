package repository

import (
	"encoding/json"
	"fmt"
	"os"
	endpoint "stat_by_sites/domain/endpoint"
	"strconv"
	"strings"
	"time"
)

type ResourceSchema struct {
	Url                  string            `json:"url"`
	Interval             string            `json:"interval"`
	Method               string            `json:"method"`
	Headers              map[string]string `json:"headers"`
	Body                 string            `json:"body"`
	ExpectedStatus       *int              `json:"expected_status"`
	ExpectedBodyContains string            `json:"expected_body_contains"`
	Timeout              string            `json:"timeout"`
	FollowRedirects      *bool             `json:"follow_redirects"`
}

type resourcesDataSchema struct {
	Resources []ResourceSchema `json:"resources"`
}

func Fetch(path string) ([]endpoint.EndpointConfig, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var schema resourcesDataSchema
	err = json.Unmarshal(data, &schema)

	if err != nil {
		return nil, err
	}

	if len(schema.Resources) == 0 {
		return nil, fmt.Errorf("config %q: resources list is empty", path)
	}

	endpoints := make([]endpoint.EndpointConfig, len(schema.Resources))
	for i, el := range schema.Resources {
		url := strings.TrimSpace(el.Url)
		if url == "" {
			return nil, fmt.Errorf("config %q: resources[%d].url is required", path, i)
		}

		parsedInterval, err := parseInterval(el.Interval)
		if err != nil {
			return nil, fmt.Errorf("config %q: resources[%d].interval is invalid: %w", path, i, err)
		}

		cfg := endpoint.NewDefaultEndpointConfig(url, parsedInterval)
		if strings.TrimSpace(el.Method) != "" {
			cfg.Method = el.Method
		}
		if len(el.Headers) > 0 {
			cfg.Headers = el.Headers
		}
		if el.Body != "" {
			cfg.Body = el.Body
		}
		if el.ExpectedStatus != nil {
			cfg.ExpectedStatus = *el.ExpectedStatus
		}
		if strings.TrimSpace(el.ExpectedBodyContains) != "" {
			cfg.ExpectedBodyContains = el.ExpectedBodyContains
		}
		if strings.TrimSpace(el.Timeout) != "" {
			timeout, err := parseInterval(el.Timeout)
			if err != nil {
				return nil, fmt.Errorf("config %q: resources[%d].timeout is invalid: %w", path, i, err)
			}
			cfg.Timeout = timeout
		}
		if el.FollowRedirects != nil {
			cfg.FollowRedirects = *el.FollowRedirects
		}

		cfg, err = endpoint.NormalizeEndpointConfig(cfg)
		if err != nil {
			return nil, fmt.Errorf("config %q: resources[%d] is invalid: %w", path, i, err)
		}

		endpoints[i] = cfg
	}

	return endpoints, nil
}

func parseInterval(raw string) (time.Duration, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return 0, fmt.Errorf("must not be empty")
	}

	if d, err := time.ParseDuration(trimmed); err == nil {
		if d <= 0 {
			return 0, fmt.Errorf("must be > 0")
		}
		return d, nil
	}

	seconds, err := strconv.Atoi(trimmed)
	if err != nil || seconds <= 0 {
		return 0, fmt.Errorf("expected duration (e.g. 5s, 1m) or positive integer seconds")
	}

	return time.Duration(seconds) * time.Second, nil
}
