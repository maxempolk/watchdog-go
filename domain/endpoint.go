package domain

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

type HealthCheckResult struct {
  URL         string
  Status      int
  Latency     time.Duration
  Error       error
}


const TrendSize = 4

func NewEndpoint(url string, interval time.Duration) Endpoint{
  return Endpoint{
    URL: url,
    Interval: interval,
    Trend: make([]bool, TrendSize),
  }
}