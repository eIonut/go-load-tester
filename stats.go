package main

import (
	"fmt"
	"time"
)

func aggregateStatistics(statistics []LoadTestStatistics) LoadTestSummary {
	summary := LoadTestSummary{
		totalRequests: len(statistics),
		errorCounts:   make(map[string]int),
	}

	var totalLatency time.Duration
	var minLatency time.Duration
	var maxLatency time.Duration

	for _, s := range statistics {
		if s.err != nil {
			summary.errorCounts[s.err.Error()]++
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

	fmt.Println()
	fmt.Println("Requests:")
	fmt.Println()
	fmt.Printf("%-22s %d\n", "Total requests:", summary.totalRequests)
	fmt.Printf("%-22s %s%d%s\n", "Successful:", green, summary.successRequests, reset)
	fmt.Printf("%-22s %s%d%s\n", "Failed:", red, summary.failedRequests, reset)

	fmt.Println()
	fmt.Println("Latency:")
	fmt.Println()
	fmt.Printf("%-22s %s\n", "Total test duration:", formatDuration(summary.totalTestDuration))
	fmt.Printf("%-22s %s\n", "Average latency:", formatDuration(summary.averageLatency))
	fmt.Printf("%-22s %s\n", "Min latency:", formatDuration(summary.minimumLatency))
	fmt.Printf("%-22s %s\n", "Max latency:", formatDuration(summary.maximumLatency))

	fmt.Println()
	fmt.Println("Throughput:")
	fmt.Println()

	fmt.Printf(
		"%-22s %s%.0f%s\n",
		"Successful Requests/sec:",
		green,
		summary.successfulRequestsPerSecond,
		reset,
	)

	fmt.Printf(
		"%-22s %s%.0f%s\n",
		"Failed Requests/sec:",
		red,
		summary.failedRequestsPerSecond,
		reset,
	)

	if len(summary.errorCounts) > 0 {
		fmt.Println()
		fmt.Println("Errors:")

		for message, count := range summary.errorCounts {
			fmt.Printf("  %s%d%s x %s\n", red, count, reset, message)
		}
	}
}
