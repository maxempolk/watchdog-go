package userlog

import "time"

type Level int

const (
	LevelInfo Level = iota
	LevelWarn
	LevelCritical
)

type Log struct{
	URL string
	Time time.Time
	Level Level
	Message string
}

func NewLog(time time.Time, level Level, message string, url string) *Log{
	return &Log{
		Time: time,
		Level: level,
		Message: message,
		URL: url,
	}
}

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelCritical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}