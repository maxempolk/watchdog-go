package formating

import (
	"fmt"
	"time"
)

func FormatTimeAgo(t time.Time) string {
	diff := time.Since(t)

	switch {
	case diff < time.Second:
		return "just now"
	case diff < time.Minute:
		return fmt.Sprintf("%ds ago", int(diff.Seconds()))
	case diff < time.Hour:
		return fmt.Sprintf("%dm ago", int(diff.Minutes()))
	default:
		return fmt.Sprintf("%dh ago", int(diff.Hours()))
	}
}

func FormatDurationMs(d time.Duration) string {
	ms := d.Milliseconds()
	return fmt.Sprintf("%d ms", ms)
}

func FormatStatusCode(status int) string {
	return fmt.Sprintf("[ %d ]", status)
}
