package logs

import "stat_by_sites/ui/components/base"

type Logs struct{
	base.Component
	IsDisplayed bool
}

func NewLogs(width int, isDisplayed bool) *Logs{
	return &Logs{
		base.Component{Width: width},
		isDisplayed,
	}
}

func (l *Logs) ToggleDisplay(){
	l.IsDisplayed = !l.IsDisplayed
}

func (l *Logs) View() string{
	return `LOGS (Recent Events)
[23:14:02] CRITICAL: db-cluster-internal returned 503 (Service Unavailable)
[23:14:10] INFO: Initializing health check for payment.gateway.net
[23:15:01] SUCCESS: api.service.com responded in 42ms`
}