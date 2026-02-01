package ui

import (
	tea "github.com/charmbracelet/bubbletea"
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
		}
	case RefreshRequestMsg:
		return m, tea.Batch(
			checkHealthCmd(*m.Service, msg.URL),
			tickEvery(msg.Interval, msg.URL), // ПЕРЕЗАПУСК таймера
		)
	case UIRefreshRequestMsg:
		endpoints := m.Service.ListEndpoints()
		m.Table.Update( endpoints )
		m.StatisticsBar.Update( endpoints )
		return m, tickEveryUI(UI_REFRESH_INTERVAL)
	case checkResultMsg:
		m.Service.UpdateEndpoint(msg.url, msg.status, msg.latency)
		m.Table.Update( m.Service.ListEndpoints() )
	}

	return m, nil
}