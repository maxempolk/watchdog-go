package commands

type CommandID string

const (
	Quit     CommandID = "quit"
	Refresh  CommandID = "refresh"
	State    CommandID = "state"
	ViewLogs CommandID = "view_logs"
	Freeze   CommandID = "freeze_logs"
)

type Command struct {
	ID          CommandID
	Key         string
	Description string
}

type Handler func() error
