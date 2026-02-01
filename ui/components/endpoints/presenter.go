package endpoints

import (
	"fmt"
	"stat_by_sites/domain"
	"time"
)

type EndpointPresenter struct{}

func FormatDurationMs(d time.Duration) string {
	ms := d.Milliseconds()
	return fmt.Sprintf("%d ms", ms)
}

func FormatTimeAgo(t time.Time) string {
	diff := time.Since(t)

	switch {
	case diff < time.Second:
		return "just now"
	case diff < time.Minute:
		return fmt.Sprintf("%ds ago", int(diff.Seconds()))
	case diff < time.Hour:
		return fmt.Sprintf("%dm ago", int(diff.Minutes()))
	default:
		return fmt.Sprintf("%dh ago", int(diff.Hours()))
	}
}

func (p EndpointPresenter) Present(list []domain.Endpoint) []Endpoint {
	out := make([]Endpoint, len(list))
	for i, e := range list {
		out[i] = Endpoint{
			URL:       e.URL,
			Status:    e.Status,
			Latency:   FormatDurationMs(e.Latency),
			LastCheck: FormatTimeAgo(e.LastCheck),
			Trend:     e.Trend,
		}
	}
	return out
}