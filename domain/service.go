package domain

import (
	"net/http"
	"time"
)

type EndpointService struct {
	repo EndpointRepository
}

func NewEndpointService(r EndpointRepository) *EndpointService {
	return &EndpointService{repo: r}
}

func (es *EndpointService) ListEndpoints() []Endpoint{
	repoKeys := es.repo.GetKeys()
	endpoints := make([]Endpoint, 0, len(repoKeys))
	for _, key := range repoKeys{
		ep, _ := es.repo.Get(key)
		endpoints = append(endpoints, ep)
	}
	return endpoints
}

func (es *EndpointService) UpdateEndpoint(url string, status int, latency time.Duration) error{
	return es.repo.Update(url, status, latency)
}

func (es *EndpointService) CheckHealth(rawUrl string) (HealthCheckResult, error){
	validatedUrl, err := validateURL(rawUrl)
	if err != nil {
		return HealthCheckResult{URL: rawUrl}, err;
	}

	client := &http.Client{Timeout: 3 * time.Second}
	start := time.Now()
	
	resp, err := client.Head(validatedUrl)
	duration := time.Since(start)

	if err != nil {
		return HealthCheckResult{URL: rawUrl, Latency: duration}, err
	}
	defer resp.Body.Close()

	return HealthCheckResult{
		URL: rawUrl,
		Latency: duration,
		Status: resp.StatusCode,
	}, nil
}

// func (es *EndpointService) UpdateEndpoint()

// func (es *EndpointService) RefreshEndpoints(){
// 	ch := make(chan HealthCheckResult)

// 	for _, k := range es.repo.GetKeys(){
// 		if endpoint, err := es.repo.Get(k); !err{
// 			go func(){
// 				ch <- CheckHealth(endpoint.URL)
// 			}()
			
// 			es.repo.Update(k, Endpoint{})
// 		}
// 	}
// }