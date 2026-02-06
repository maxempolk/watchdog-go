package endpoint

import (
	"time"
)

type EndpointService struct {
	repo EndpointRepository
	healthChecker HealthChecker
}

func NewEndpointService(r EndpointRepository, hc HealthChecker) *EndpointService {
	return &EndpointService{
		repo: r,
		healthChecker: hc,
	}
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

func (es *EndpointService) CheckHealth(url string) (HealthCheckResult, error){
	return es.healthChecker.Check(url)
}