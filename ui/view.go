package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)


func (m model) View() string {
	EMPTY := ""
	LINE := strings.Repeat("─", BASE_WIDTH)
	
	views := make([]string, 0, 8)

	views = append(views, m.Header.View())
	views = append(views, m.StatisticsBar.View())
	views = append(views, LINE)
	views = append(views, EMPTY)
	views = append(views, m.Table.View())
	views = append(views, LINE)
	if !m.Logs.IsDisplayed{
		views = append(views, EMPTY)
		views = append(views, m.Logs.View())
		views = append(views, EMPTY)
		views = append(views, LINE)
	}
	views = append(views, m.CommandsHelper.View())

	return lipgloss.JoinVertical(
		lipgloss.Left,
		views...,
	)
}