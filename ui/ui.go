package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const BASE_WIDTH = 80
const REFRESH_INTERVAL = 15 * time.Second
const UI_REFRESH_INTERVAL = 500 * time.Millisecond

type RefreshRequestMsg struct{
	Interval time.Duration
	URL string
}

func tickEvery(duration time.Duration, url string) tea.Cmd {
	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return RefreshRequestMsg{
			Interval: duration,
			URL: url,
		}
	})
}

type UIRefreshRequestMsg time.Time

func tickEveryUI(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return UIRefreshRequestMsg(t)
	})
}