package commands

import "fmt"

type Registry map[CommandID]Handler

type Manager struct {
	handlers Registry
}

func NewManager() *Manager {
	return &Manager{
		handlers: make(Registry),
	}
}

func (m *Manager) Register(id CommandID, handler Handler) {
	m.handlers[id] = handler
}

func (m *Manager) Execute(id CommandID) error {
	handler, exists := m.handlers[id]
	if !exists {
		return fmt.Errorf("unknown command: %s", id)
	}
	return handler()
}

// Вспомогательная функция: переводит клавишу в CommandID
func KeyToCommand(key string) CommandID {
	switch key {
	case "q":
		return Quit
	case "r":
		return Refresh
	case "s":
		return State
	case "l":
		return ViewLogs
	case "f":
		return Freeze
	}
	return ""
}

func GetAllCommands() []Command {
	return []Command{
		{Quit, "Q", "Quit"},
		{Refresh, "R", "Refresh"},
		{State, "S", "State"},
		{ViewLogs, "L", "View Logs"},
		{Freeze, "F", "Freeze/Unfreeze Logs"},
	}
}
