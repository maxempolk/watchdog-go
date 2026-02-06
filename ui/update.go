package ui

import (
	"fmt"
	"log/slog"
	"stat_by_sites/domain/endpoint"
	userlog "stat_by_sites/domain/log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type checkResultMsg struct {
	url     string
	status  int
	latency time.Duration
	err     error
}

func checkHealthCmd(
	svc endpoint.EndpointService,
	url string,
) tea.Cmd {
	return func() tea.Msg {
		healthCheckResult, err := svc.CheckHealth(url)

		if err != nil {
			slog.Error("check health failed", "url", url, "error", err)
		}

		return checkResultMsg{
			url:     healthCheckResult.URL,
			status:  healthCheckResult.Status,
			latency: healthCheckResult.Latency,
			err:     err,
		}
	}
}

func (m *model) appendCheckResultLog(msg checkResultMsg) {
	if m.LogService == nil {
		return
	}

	m.LogService.Add(*userlog.NewLog(
		time.Now(),
		resolveLogLevel(msg),
		buildCheckLogMessage(msg),
		msg.url,
	))
}

func (m *model) refreshLogs() {
	if m.LogService == nil || m.Logs == nil {
		return
	}

	total := m.LogService.Count()
	list, err := m.LogService.ListRecent(logsListLimit)
	if err != nil {
		slog.Error("failed to refresh logs", "error", err)
		return
	}

	m.Logs.Update(list, total)
}

func resolveLogLevel(msg checkResultMsg) userlog.Level {
	if msg.err != nil || msg.status >= 500 || msg.status == 0 {
		return userlog.LevelCritical
	}

	if msg.status >= 400 {
		return userlog.LevelWarn
	}

	return userlog.LevelInfo
}

func buildCheckLogMessage(msg checkResultMsg) string {
	if msg.err != nil {
		return fmt.Sprintf("%s check failed: %v", msg.url, msg.err)
	}

	return fmt.Sprintf("%s returned HTTP %d in %dms", msg.url, msg.status, msg.latency.Milliseconds())
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
