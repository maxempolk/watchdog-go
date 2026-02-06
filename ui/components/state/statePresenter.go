package state

import (
	"sort"
	"stat_by_sites/domain/endpoint"
	"stat_by_sites/ui/formating"
)

type EndpointState struct {
	URL       string
	Status    string
	Latency   string
	Result    string
	LastCheck string
}

type StatePresenter struct{}

func (p StatePresenter) Present(list []endpoint.Endpoint) []EndpointState {
	sorted := make([]endpoint.Endpoint, len(list))
	copy(sorted, list)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].URL < sorted[j].URL
	})

	out := make([]EndpointState, len(sorted))
	for i, e := range sorted {
		out[i] = EndpointState{
			URL:       e.URL,
			Status:    formatStatus(e),
			Latency:   formatLatency(e),
			Result:    formatResult(e),
			LastCheck: formatLastCheck(e),
		}
	}

	return out
}

func formatStatus(e endpoint.Endpoint) string {
	if e.LastCheck.IsZero() {
		return "-"
	}

	return formating.FormatStatusCode(e.Status)
}

func formatLatency(e endpoint.Endpoint) string {
	if e.LastCheck.IsZero() {
		return "-"
	}

	return formating.FormatDurationMs(e.Latency)
}

func formatResult(e endpoint.Endpoint) string {
	if e.LastCheck.IsZero() {
		return "NO DATA"
	}

	if e.IsHealthy() {
		return "SUCCESS"
	}

	return "FAILED"
}

func formatLastCheck(e endpoint.Endpoint) string {
	if e.LastCheck.IsZero() {
		return "-"
	}

	return formating.FormatTimeAgo(e.LastCheck)
}
