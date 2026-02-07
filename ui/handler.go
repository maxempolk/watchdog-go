package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"log/slog"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "r":
			// Ручное обновление по нажатию R
			return m, func() tea.Msg {
				return UIRefreshRequestMsg{}
			}
		case "l":
			m.Logs.ToggleDisplay()
			return m, nil
		case "s":
			m.State.ToggleDisplay()
			return m, nil
		case "f":
			total := 0
			if m.LogService != nil {
				total = m.LogService.Count()
			}
			m.Logs.TogglePause(total)
			m.refreshLogs()
			return m, nil
		case "up", "k":
			if m.Logs.IsDisplayed {
				m.Logs.ScrollUp()
			}
			return m, nil
		case "down", "j":
			if m.Logs.IsDisplayed {
				m.Logs.ScrollDown()
			}
			return m, nil
		}
	case RefreshRequestMsg:
		return m, tea.Batch(
			checkHealthCmd(*m.Service, msg.URL),
			tickEvery(msg.Interval, msg.URL), // ПЕРЕЗАПУСК таймера
		)
	case UIRefreshRequestMsg:
		endpoints := m.Service.ListEndpoints()
		m.Table.Update(endpoints)
		m.StatisticsBar.Update(endpoints)
		m.State.Update(endpoints)
		m.refreshLogs()
		return m, tickEveryUI(UI_REFRESH_INTERVAL)
	case checkResultMsg:
		status := msg.status
		if msg.err != nil {
			status = 0
		}

		if err := m.Service.UpdateEndpoint(msg.url, status, msg.latency); err != nil {
			slog.Error("failed to update endpoint", "url", msg.url, "error", err)
			return m, nil
		}
		m.appendCheckResultLog(msg)
		endpoints := m.Service.ListEndpoints()
		m.Table.Update(endpoints)
		m.State.Update(endpoints)
		m.refreshLogs()
	}

	return m, nil
}
