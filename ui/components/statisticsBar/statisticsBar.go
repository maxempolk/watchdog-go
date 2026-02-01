package statisticsBar

import (
	"fmt"
	"stat_by_sites/domain"
	"stat_by_sites/ui/components/base"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

type StatisticsBar struct{
	base.Component
	uptime float64
	errors int
	latency string
}

func NewStatisticsBar(width int) *StatisticsBar{
	return &StatisticsBar{ 
		Component: base.Component{Width: width},
	}
}

func (sb *StatisticsBar) Update(endpoints []domain.Endpoint){
	stats := CalculateStats(endpoints)

	sb.uptime = float64(stats.Healthy) / float64(stats.Total)
	sb.errors = stats.Errors
	sb.latency = fmt.Sprintf("%dms", stats.AvgLatency)
}

func (sb *StatisticsBar) View() string{
	var (
		labelStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Bold(true)
		errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
		latencyStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	)
	
	pg := progress.New(
		progress.WithDefaultGradient(),
		progress.WithoutPercentage(), // Мы напишем проценты сами
		progress.WithWidth(20),
	)

	// 2. Формируем части строки
	upSection := fmt.Sprintf("%s [%s] %.0f%%", 
		labelStyle.Render("UP"), 
		pg.ViewAs(sb.uptime), 
		sb.uptime*100,
	)

	errSection := fmt.Sprintf("%s %s %d", 
		labelStyle.Render("ERRORS"), 
		errorStyle.Render("[!]"), 
		sb.errors,
	)

	latencySection := fmt.Sprintf("%s %s %s", 
		labelStyle.Render("AVG LATENCY"), 
		latencyStyle.Render("[~]"),  
		sb.latency,
	)

	// 3. Соединяем секции с отступами
	content := fmt.Sprintf("%s    %s    %s", upSection, errSection, latencySection)

	// 4. ВЫРАВНИВАНИЕ ПО ЦЕНТРУ
	// lipgloss.Place берет контент и центрирует его в пространстве заданной ширины
	return lipgloss.Place(
		sb.Width,           // На какую ширину центрировать
		1,               // Высота контейнера
		lipgloss.Center, // Горизонтальное выравнивание
		lipgloss.Center, // Вертикальное выравнивание
		content,
	)
}