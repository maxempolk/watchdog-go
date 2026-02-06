package healthcheck

import (
	"net/http"
	"stat_by_sites/domain/endpoint"
	"stat_by_sites/domain/shared"
	"time"
)

type HealthChecker struct{}

func (hc *HealthChecker) Check(url string) (endpoint.HealthCheckResult, error){
	validatedUrl, err := shared.ValidateURL(url)
	if err != nil {
		return endpoint.HealthCheckResult{URL: url}, err;
	}

	client := &http.Client{Timeout: 3 * time.Second}
	start := time.Now()
	
	resp, err := client.Head(validatedUrl)
	duration := time.Since(start)

	if err != nil {
		return endpoint.HealthCheckResult{URL: url, Latency: duration}, err
	}
	defer resp.Body.Close()

	return endpoint.HealthCheckResult{
		URL: url,
		Latency: duration,
		Status: resp.StatusCode,
	}, nil
}