package endpoints

import (
	"stat_by_sites/domain"
	"stat_by_sites/ui/formating"
)

type EndpointPresenter struct{}



func (p EndpointPresenter) Present(list []domain.Endpoint) []Endpoint {
	out := make([]Endpoint, len(list))
	for i, e := range list {
		out[i] = Endpoint{
			URL:       e.URL,
			Status:    e.Status,
			Latency:   formating.FormatDurationMs(e.Latency),
			LastCheck: formating.FormatTimeAgo(e.LastCheck),
			Trend:     e.Trend,
		}
	}
	return out
}