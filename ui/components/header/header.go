package header

import (
	"stat_by_sites/ui/components/base"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type Header struct{
	base.Component
}

func NewHeader(width int) *Header {
	return &Header{ base.Component{Width: width} }
}

func (h Header) View() string {
	// 1. Определяем стиль рамки
	headerStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()). // Можно использовать NormalBorder() для острых углов
		BorderForeground(lipgloss.Color("63")). // Цвет рамки (фиолетовый)
		Padding(0, 1).                          // Отступы внутри (сверху/снизу 0, слева/справа 1)
		Width(80)                               // Жесткая ширина всей рамки

	// 2. Стили для текста внутри
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA"))
	timeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#7D7D7D"))

	// 3. Контент
	title := titleStyle.Render("HEALTH MONITOR v1.0.2")
	currTime := timeStyle.Render("Current Time: " + time.Now().Format("2006-01-02 15:04"))

	// 4. Магия выравнивания:
	// Считаем сколько места нужно между левым и правым текстом
	// Width(80) минус границы (2) и паддинги (2) = 76 символов полезного пространства
	width := 76 
	
	// Используем либу для создания "заполнителя" (пробелов) между текстами
	space := lipgloss.NewStyle().Width(width - lipgloss.Width(title) - lipgloss.Width(currTime)).Render("")

	// Собираем всё вместе
	content := lipgloss.JoinHorizontal(lipgloss.Center, title, space, currTime)
	
	return headerStyle.Render(content)
}