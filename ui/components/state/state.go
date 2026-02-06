package state

import (
	"fmt"
	"stat_by_sites/domain/endpoint"
	"stat_by_sites/ui/components/base"
	"strings"
)

const (
	Header            = "STATE (Last Status And Results)"
	endpointColWidth  = 34
	statusColWidth    = 10
	latencyColWidth   = 10
	resultColWidth    = 10
	lastCheckColWidth = 12
)

type State struct {
	base.Component
	IsDisplayed bool
	rows        []EndpointState
	presenter   StatePresenter
}

func NewState(width int, isDisplayed bool) *State {
	return &State{
		Component:   base.Component{Width: width},
		IsDisplayed: isDisplayed,
		rows:        make([]EndpointState, 0),
		presenter:   StatePresenter{},
	}
}

func (s *State) ToggleDisplay() {
	s.IsDisplayed = !s.IsDisplayed
}

func (s *State) Update(endpoints []endpoint.Endpoint) {
	s.rows = s.presenter.Present(endpoints)
}

func (s State) View() string {
	if len(s.rows) == 0 {
		return fmt.Sprintf("%s\nThere are no endpoint states...", Header)
	}

	lines := make([]string, 0, len(s.rows)+3)
	lines = append(lines, Header)
	lines = append(lines, formatStateRow("ENDPOINT", "STATUS", "LATENCY", "RESULT", "LAST CHECK"))
	lines = append(lines, strings.Repeat("─", s.Width))

	for _, row := range s.rows {
		lines = append(lines, formatStateRow(row.URL, row.Status, row.Latency, row.Result, row.LastCheck))
	}

	return strings.Join(lines, "\n")
}

func formatStateRow(endpoint, status, latency, result, lastCheck string) string {
	return fmt.Sprintf(
		"%-*s %-*s %-*s %-*s %-*s",
		endpointColWidth, truncateStateValue(endpoint, endpointColWidth),
		statusColWidth, truncateStateValue(status, statusColWidth),
		latencyColWidth, truncateStateValue(latency, latencyColWidth),
		resultColWidth, truncateStateValue(result, resultColWidth),
		lastCheckColWidth, truncateStateValue(lastCheck, lastCheckColWidth),
	)
}

func truncateStateValue(value string, maxLength int) string {
	if maxLength <= 3 || len(value) <= maxLength {
		return value
	}

	return value[:maxLength-3] + "..."
}
