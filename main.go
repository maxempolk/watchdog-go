package main

import (
	"log"
	"log/slog"
	"stat_by_sites/domain"
	"stat_by_sites/internal/config"
	"stat_by_sites/internal/logger"
	"stat_by_sites/internal/repository"
	"stat_by_sites/ui"

	tea "github.com/charmbracelet/bubbletea"
)



func main(){
	if err := logger.Init(); err != nil {
		panic(err)
	}
	slog.Info("app started")

	cfg := config.NewConfig()
	memoryEndpointRepository := repository.NewMemoryEndpointRepository()
	endpointService := domain.NewEndpointService(memoryEndpointRepository)

	if err := cfg.Validate(); err != nil {
		log.Fatal(err) // Выведет: флаг -file является обязательным
	}
	
	endpointsConfig, err := repository.Fetch(cfg.FilePath)

	if err != nil{
		log.Fatal(err)
	}
	
	memoryEndpointRepository.Add( endpointsConfig... )

	m := ui.InitialModel(endpointService)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}