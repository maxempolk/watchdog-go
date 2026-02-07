package main

import (
	"log"
	"log/slog"
	"os"
	"stat_by_sites/domain/endpoint"
	userlog "stat_by_sites/domain/log"
	"stat_by_sites/internal/config"
	"stat_by_sites/internal/healthcheck"
	"stat_by_sites/internal/logger"
	"stat_by_sites/internal/repository"
	"stat_by_sites/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if err := logger.Init(); err != nil {
		panic(err)
	}
	slog.Info("app started")

	runtimeCfg, err := config.ParseRuntimeConfig(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	memoryEndpointRepository := repository.NewMemoryEndpointRepository()
	memoryLogRepository := repository.NewMemoryLogRepository()
	endpointService := endpoint.NewEndpointService(
		memoryEndpointRepository,
		&healthcheck.HealthChecker{},
	)
	logService := userlog.NewLogService(memoryLogRepository)

	var endpointsConfig []endpoint.EndpointConfig
	switch runtimeCfg.Mode {
	case config.ModeDefault, config.ModeFile:
		endpointsConfig, err = repository.Fetch(runtimeCfg.FilePath)
		if err != nil {
			log.Fatal(err)
		}
	case config.ModeSites:
		endpointsConfig = runtimeCfg.Endpoints
	default:
		log.Fatalf("unsupported runtime mode: %s", runtimeCfg.Mode)
	}

	if err := memoryEndpointRepository.Add(endpointsConfig...); err != nil {
		log.Fatal(err)
	}

	m := ui.InitialModel(endpointService, logService)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
