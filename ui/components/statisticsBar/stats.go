package statisticsBar

import "stat_by_sites/domain"

type Stats struct {
	Total        int
	Healthy      int
	Errors       int
	AvgLatency 	 int
}

func CalculateStats(endpoints []domain.Endpoint) Stats {
	var healthy int
	var latencySum int

	for _, e := range endpoints {
		if e.IsHealthy() {
			healthy++
			latencySum += int(e.Latency.Milliseconds())
		}
	}

	avg := 0
	if healthy > 0 {
		avg = latencySum / healthy
	}

	return Stats{
		Total:        len(endpoints),
		Healthy:      healthy,
		Errors:       len(endpoints) - healthy,
		AvgLatency: 	avg,
	}
}