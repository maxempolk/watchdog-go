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
	Url      string `json:"url"`
	Interval string `json:"interval"`
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

		endpoints[i] = endpoint.EndpointConfig{
			URL:      url,
			Interval: parsedInterval,
		}
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
