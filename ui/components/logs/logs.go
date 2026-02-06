package logs

import (
	"fmt"
	userlog "stat_by_sites/domain/log"
	"stat_by_sites/ui/components/base"
	"strings"
)

const Header = "LOGS (Recent Events)"
const visibleRows = 6

type Logs struct {
	base.Component
	IsDisplayed bool
	logs        []userlog.Log
	scroll      int
	IsPaused    bool
	newEntries  int
	frozenTotal int
	*LogPresenter
}

func NewLogs(width int, isDisplayed bool) *Logs {
	return &Logs{
		base.Component{Width: width},
		isDisplayed,
		make([]userlog.Log, 0),
		0,
		false,
		0,
		0,
		&LogPresenter{},
	}
}

func (l *Logs) ToggleDisplay() {
	l.IsDisplayed = !l.IsDisplayed
}

func (l *Logs) View() string {
	header := Header
	if l.IsPaused {
		header = fmt.Sprintf("%s ⏸ PAUSED — %d new entries", Header, l.newEntries)
	}

	count := len(l.logs)
	if count == 0 {
		return fmt.Sprintf("%s\nThere are no logs...", header)
	}

	start := l.scroll
	end := start + visibleRows
	if end > count {
		end = count
	}

	rows := make([]string, 0, end-start+2)
	rows = append(rows, header)

	for _, v := range l.logs[start:end] {
		rows = append(rows, l.Present(v))
	}

	if count > visibleRows {
		rows = append(rows, fmt.Sprintf("Showing %d-%d of %d (Up/Down to scroll)", start+1, end, count))
	}

	return strings.Join(rows, "\n")
	//	return fmt.Sprintf(`%s\n
	//
	// [23:14:02] CRITICAL: db-cluster-internal returned 503 (Service Unavailable)
	// [23:14:10] INFO: Initializing health check for payment.gateway.net
	// [23:15:01] SUCCESS: api.service.com responded in 42ms`, Header)
}

func (l *Logs) Update(logs []userlog.Log, total int) {
	if l.IsPaused {
		if total > l.frozenTotal {
			l.newEntries = total - l.frozenTotal
		} else {
			l.newEntries = 0
		}
		return
	}

	updated := make([]userlog.Log, len(logs))
	copy(updated, logs)
	l.logs = updated
	l.newEntries = 0
	l.frozenTotal = total
	l.ensureScrollBounds()
}

func (l *Logs) TogglePause(total int) {
	l.IsPaused = !l.IsPaused
	if l.IsPaused {
		l.frozenTotal = total
		l.newEntries = 0
		return
	}

	l.newEntries = 0
}

func (l *Logs) ScrollDown() {
	l.scroll++
	l.ensureScrollBounds()
}

func (l *Logs) ScrollUp() {
	l.scroll--
	l.ensureScrollBounds()
}

func (l *Logs) ensureScrollBounds() {
	if l.scroll < 0 {
		l.scroll = 0
		return
	}

	maxScroll := len(l.logs) - visibleRows
	if maxScroll < 0 {
		maxScroll = 0
	}

	if l.scroll > maxScroll {
		l.scroll = maxScroll
	}
}
