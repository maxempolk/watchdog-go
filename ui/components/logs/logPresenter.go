package logs

import (
	"fmt"
	userlog "stat_by_sites/domain/log"
)

type LogPresenter struct{}

func (lp LogPresenter) Present(log userlog.Log) string {
	formattedTime := log.Time.Format("15:04:05")

	return fmt.Sprintf(`[%s] %s: %s`, formattedTime, log.Level.String(), log.Message)
}

// [23:14:02] CRITICAL: db-cluster-internal returned 503 (Service Unavailable)
// [23:14:10] INFO: Initializing health check for payment.gateway.net
// [23:15:01] SUCCESS: api.service.com responded in 42ms
