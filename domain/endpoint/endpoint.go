package endpoint

import (
	"time"
)

type EndpointConfig struct{
  URL      string
  Interval time.Duration
}

type Endpoint struct {
  URL       string
  Status    int
  Latency   time.Duration
  LastCheck time.Time
  Trend     []bool
  Interval  time.Duration
}

const TrendSize = 4

func NewEndpoint(url string, interval time.Duration) Endpoint{
  return Endpoint{
    URL: url,
    Interval: interval,
    Trend: make([]bool, TrendSize),
  }
}

func (e Endpoint) IsHealthy() bool{
  return e.Status >= 200 && e.Status < 400
}



type HealthCheckResult struct {
  URL         string
  Status      int
  Latency     time.Duration
}

type HealthChecker interface { 
  Check(url string) (HealthCheckResult, error) 
}