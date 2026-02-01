package commands

type CommandID string

const (
    Quit    CommandID = "quit"
    Refresh CommandID = "refresh"
    Settings CommandID = "settings"
    ViewLogs CommandID = "view_logs"
)

type Command struct {
    ID          CommandID  
    Key         string
    Description string
}

type Handler func() error