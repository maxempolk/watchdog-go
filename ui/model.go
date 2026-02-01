package ui

import (
	"fmt"
	"log/slog"
	"stat_by_sites/domain"
	"stat_by_sites/internal/commands"
	commandshelper "stat_by_sites/ui/components/commandsHelper"
	"stat_by_sites/ui/components/endpoints"
	"stat_by_sites/ui/components/header"
	"stat_by_sites/ui/components/logs"
	"stat_by_sites/ui/components/statisticsBar"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Service			   *domain.EndpointService
	Header         *header.Header
	StatisticsBar  *statisticsBar.StatisticsBar
	Table          *endpoints.Table
	Logs           *logs.Logs
	CommandsHelper *commandshelper.CommandHelper
}

func (m model) Init() tea.Cmd {
	// TODO: проверить точно ли в момент запуска интерфейса listEnpoints уже не пустой
	cmds := make([]tea.Cmd, 0, len(m.Service.ListEndpoints())*2+2)
	cmds = append(cmds, 
		tea.Cmd(tea.ClearScreen), 
		tickEveryUI(UI_REFRESH_INTERVAL),
	)
	for _, e := range m.Service.ListEndpoints() {
		cmds = append(cmds, checkHealthCmd(*m.Service, e.URL))
	}
	
	slog.Info("UI init cmds: " + fmt.Sprint(cmds))
	for _, v := range m.Service.ListEndpoints() {
		cmds = append(cmds, tickEvery(v.Interval, v.URL))
	}

	return tea.Batch(cmds...)
}

func InitialModel(service *domain.EndpointService) *model {
	table := endpoints.NewTable(BASE_WIDTH)
	table.Update(service.ListEndpoints())
	
	return &model{
		Service: service,
		Header: header.NewHeader(BASE_WIDTH),
		StatisticsBar: statisticsBar.NewStatisticsBar(BASE_WIDTH),
		Table: table,
		Logs: logs.NewLogs(BASE_WIDTH, false),
		CommandsHelper: commandshelper.NewCommandHelper(
				BASE_WIDTH,
				commands.GetAllCommands(),
		),
	}
}