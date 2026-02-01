package ui

import (
	"log/slog"

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
		// Периодическое обновление каждые REFRESH_INTERVAL секунд
		// Проверяем все эндпоинты в фоне
		// cmds := make([]tea.Cmd, 0, len(m.Table.Data))
		// for _, e := range m.Table.Data {
		// 	cmds = append(cmds, checkHealthCmd(*m.Service, e.URL))
		// }
		slog.Info( "tickMsg: " + msg.URL )
		return m, tea.Batch(
			checkHealthCmd(*m.Service, msg.URL),
			tickEvery(msg.Interval, msg.URL), // ПЕРЕЗАПУСК таймера
		)
	case UIRefreshRequestMsg:
		m.Table.Update( m.Service.ListEndpoints() )
		return m, tickEveryUI(UI_REFRESH_INTERVAL)
	case checkResultMsg:
		m.Service.UpdateEndpoint(msg.url, msg.status, msg.latency)
		m.Table.Update( m.Service.ListEndpoints() )
	}

	return m, nil
}