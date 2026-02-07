package healthcheck

import (
	"fmt"
	"io"
	"net/http"
	"stat_by_sites/domain/endpoint"
	"stat_by_sites/domain/shared"
	"strings"
	"time"
)

type HealthChecker struct{}

func (hc *HealthChecker) Check(ep endpoint.Endpoint) (endpoint.HealthCheckResult, error) {
	validatedURL, err := shared.ValidateURL(ep.URL)
	if err != nil {
		return endpoint.HealthCheckResult{URL: ep.URL}, err
	}

	req, err := buildRequest(ep, validatedURL)
	if err != nil {
		return endpoint.HealthCheckResult{URL: ep.URL}, err
	}

	client := &http.Client{Timeout: ep.Timeout}
	if !ep.FollowRedirects {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	start := time.Now()
	resp, err := client.Do(req)
	duration := time.Since(start)

	if err != nil {
		return endpoint.HealthCheckResult{URL: ep.URL, Latency: duration}, err
	}
	defer resp.Body.Close()

	result := endpoint.HealthCheckResult{
		URL:     ep.URL,
		Latency: duration,
		Status:  resp.StatusCode,
	}

	if err := validateExpectations(ep, resp); err != nil {
		return result, err
	}

	return result, nil
}

func buildRequest(ep endpoint.Endpoint, validatedURL string) (*http.Request, error) {
	var bodyReader io.Reader
	if ep.Method != endpoint.MethodHead && ep.Body != "" {
		bodyReader = strings.NewReader(ep.Body)
	}

	req, err := http.NewRequest(ep.Method, validatedURL, bodyReader)
	if err != nil {
		return nil, err
	}

	for k, v := range ep.Headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

func validateExpectations(ep endpoint.Endpoint, resp *http.Response) error {
	if ep.ExpectedStatus != 0 {
		if resp.StatusCode != ep.ExpectedStatus {
			return fmt.Errorf("expected status %d, got %d", ep.ExpectedStatus, resp.StatusCode)
		}
	} else if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	if ep.ExpectedBodyContains == "" {
		return nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body for expectation check: %w", err)
	}

	if !strings.Contains(string(bodyBytes), ep.ExpectedBodyContains) {
		return fmt.Errorf("response body does not contain expected substring")
	}

	return nil
}
