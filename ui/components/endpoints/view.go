package endpoints

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (e Table) View() string{
	var rows []string

	header := lipgloss.JoinHorizontal(lipgloss.Top, e.buildHeaders())
	rows = append(rows, header)
	rows = append(rows, strings.Repeat("─", e.Width))

	for _, endpoint := range e.Data {
		rows = append(rows, e.buildRow(endpoint))
	}

	return strings.Join(rows, "\n")
}