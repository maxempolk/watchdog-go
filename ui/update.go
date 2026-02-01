package ui

import (
	"log/slog"
	"stat_by_sites/domain"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type checkResultMsg struct {
	url      string
	status   int
	latency  time.Duration
	err      error
}

func checkHealthCmd(
	svc domain.EndpointService,
	url string,
) tea.Cmd {
	return func() tea.Msg {
		healthCheckResult, err := svc.CheckHealth(url)

		if err != nil{
			slog.Error("CheckHealthCmd: " + err.Error())
		}

		return checkResultMsg{
			url:     healthCheckResult.URL,
			status:  healthCheckResult.Status,
			latency: healthCheckResult.Latency,
			err:     healthCheckResult.Error,
		}
	}
}

// func checkAllEndpoints(
// 	svc domain.EndpointService,
// 	table []domain.Endpoint,
// ) tea.Cmd {
// 	return func() tea.Msg {
// 		cmds := make([]tea.Cmd, len(table))
// 		for i, endpoint := range table {
// 			cmds[i] = checkHealthCmd(svc, endpoint.URL);
// 		}
// 		return tea.Batch(cmds...)()
// 	}
// }