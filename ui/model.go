package ui

import (
	"fmt"
	"log/slog"
	"stat_by_sites/domain/endpoint"
	userlog "stat_by_sites/domain/log"
	"stat_by_sites/internal/commands"
	commandshelper "stat_by_sites/ui/components/commandsHelper"
	"stat_by_sites/ui/components/endpoints"
	"stat_by_sites/ui/components/header"
	"stat_by_sites/ui/components/logs"
	"stat_by_sites/ui/components/state"
	"stat_by_sites/ui/components/statisticsBar"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Service        *endpoint.EndpointService
	LogService     *userlog.LogService
	Header         *header.Header
	StatisticsBar  *statisticsBar.StatisticsBar
	Table          *endpoints.Table
	State          *state.State
	Logs           *logs.Logs
	CommandsHelper *commandshelper.CommandHelper
}

const logsListLimit = 200

func (m model) Init() tea.Cmd {
	// TODO: проверить точно ли в момент запуска интерфейса listEnpoints уже не пустой
	cmds := make([]tea.Cmd, 0, len(m.Service.ListEndpoints())*2+2)
	cmds = append(cmds,
		tea.Cmd(tea.ClearScreen),
		tickEveryUI(UI_REFRESH_INTERVAL),
	)
	for _, e := range m.Service.ListEndpoints() {
		if m.LogService != nil {
			m.LogService.Add(*userlog.NewLog(
				time.Now(),
				userlog.LevelInfo,
				"initializing health check",
				e.URL,
			))
		}
		cmds = append(cmds, checkHealthCmd(*m.Service, e.URL))
	}

	slog.Info("UI init cmds: " + fmt.Sprint(cmds))
	for _, v := range m.Service.ListEndpoints() {
		cmds = append(cmds, tickEvery(v.Interval, v.URL))
	}

	return tea.Batch(cmds...)
}

func InitialModel(service *endpoint.EndpointService, logService *userlog.LogService) *model {
	table := endpoints.NewTable(BASE_WIDTH)
	table.Update(service.ListEndpoints())
	logsComponent := logs.NewLogs(BASE_WIDTH, false)
	stateComponent := state.NewState(BASE_WIDTH, false)
	stateComponent.Update(service.ListEndpoints())

	if logService != nil {
		total := logService.Count()
		list, err := logService.ListRecent(logsListLimit)
		if err != nil {
			slog.Error("failed to load logs", "error", err)
		} else {
			logsComponent.Update(list, total)
		}
	}

	return &model{
		Service:       service,
		LogService:    logService,
		Header:        header.NewHeader(BASE_WIDTH),
		StatisticsBar: statisticsBar.NewStatisticsBar(BASE_WIDTH),
		Table:         table,
		State:         stateComponent,
		Logs:          logsComponent,
		CommandsHelper: commandshelper.NewCommandHelper(
			BASE_WIDTH,
			commands.GetAllCommands(),
		),
	}
}
