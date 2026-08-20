package main

import (
	"fmt"
	"time"
)

func formatDuration(duration time.Duration) string {
	switch {
	case duration < time.Second:
		return fmt.Sprintf("%.2fms", duration.Seconds()*1000)

	case duration < time.Minute:
		return fmt.Sprintf("%.2fs", duration.Seconds())

	default:
		return fmt.Sprintf("%.2fmin", duration.Minutes())
	}
}
