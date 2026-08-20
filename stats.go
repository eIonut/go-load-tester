package main

import (
	"fmt"
	"time"
)

func aggregateStatistics(statistics []LoadTestStatistics) LoadTestSummary {
	summary := LoadTestSummary{
		totalRequests: len(statistics),
	}

	var totalLatency time.Duration
	var minLatency time.Duration
	var maxLatency time.Duration

	for _, s := range statistics {
		if s.err != nil {
			summary.failedRequests++
			continue
		}

		if summary.successRequests == 0 {
			minLatency = s.elapsedTime
			maxLatency = s.elapsedTime
		} else {
			if s.elapsedTime < minLatency {
				minLatency = s.elapsedTime
			}

			if s.elapsedTime > maxLatency {
				maxLatency = s.elapsedTime
			}
		}
		summary.successRequests++
		totalLatency += s.elapsedTime
	}

	if summary.successRequests > 0 {
		summary.averageLatency =
			totalLatency / time.Duration(summary.successRequests)
		summary.minimumLatency = minLatency
		summary.maximumLatency = maxLatency
	}

	return summary
}

const (
	reset  = "\033[0m"
	green  = "\033[32m"
	red    = "\033[31m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
)

func printLoadTestSummary(summary LoadTestSummary) {
	fmt.Printf("%sLoad test completed%s\n\n", cyan, reset)

	fmt.Printf("%-22s %d\n", "Total requests:", summary.totalRequests)
	fmt.Printf("%-22s %s%d%s\n", "Successful:", green, summary.successRequests, reset)
	fmt.Printf("%-22s %s%d%s\n", "Failed:", red, summary.failedRequests, reset)

	fmt.Printf("%-22s %s\n", "Total test duration:", formatDuration(summary.totalTestDuration))
	fmt.Printf("%-22s %s\n", "Average latency:", formatDuration(summary.averageLatency))
	fmt.Printf("%-22s %s\n", "Min latency:", formatDuration(summary.minimumLatency))
	fmt.Printf("%-22s %s\n", "Max latency:", formatDuration(summary.maximumLatency))

	fmt.Printf(
		"%-22s %s%.0f%s\n",
		"Requests/sec:",
		yellow,
		summary.requestsPerSecond,
		reset,
	)
}
