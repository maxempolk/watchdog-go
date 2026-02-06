package repository

import (
	"fmt"
	"log/slog"
	endpoint "stat_by_sites/domain/endpoint"
	"sync"
	"time"
)

type MemoryEndpointRepository struct {
	mu        sync.RWMutex
	endpoints map[string]endpoint.Endpoint
}

func NewMemoryEndpointRepository() *MemoryEndpointRepository{
	return &MemoryEndpointRepository{
		endpoints: make(map[string]endpoint.Endpoint),
	}
}

func (r *MemoryEndpointRepository) Add(endpoints ...endpoint.EndpointConfig) error{
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, v := range endpoints {
		r.endpoints[v.URL] = endpoint.NewEndpoint(v.URL, v.Interval)
	}
	return nil
}

func (r *MemoryEndpointRepository) Update(url string, status int, latency time.Duration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if v, c := r.endpoints[url]; c{
		v.Status = status
		v.Latency = latency
		v.LastCheck = time.Now()
		
		// TODO: нужно ли хранить новые статусы в конце тренда
		if len(v.Trend) >= endpoint.TrendSize {
			v.Trend = v.Trend[1:]
		}

		v.Trend = append(v.Trend, status >= 200 && status < 400)

		slog.Info("Endpoint update: " + fmt.Sprint(v) )

		r.endpoints[url] = v
	}

	return nil
}

func (r *MemoryEndpointRepository) Get(url string) (endpoint.Endpoint, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	v, ok := r.endpoints[url]
	return v, ok
}

func (r *MemoryEndpointRepository) GetKeys() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	keys := make([]string, 0, len(r.endpoints))
	for k := range r.endpoints {
		keys = append(keys, k)
	}

	return keys
}