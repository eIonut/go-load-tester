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

func printLoadTestSummary(summary LoadTestSummary) {
	fmt.Println("Load test completed")
	fmt.Println()
	fmt.Println("Total requests:", summary.totalRequests)
	fmt.Println("Successful:", summary.successRequests)
	fmt.Println("Failed:", summary.failedRequests)
	fmt.Println()

	fmt.Println(
		"Total test duration:",
		formatDuration(summary.totalTestDuration),
	)
	fmt.Println(
		"Average latency:",
		formatDuration(summary.averageLatency),
	)
	fmt.Println(
		"Min latency:",
		formatDuration(summary.minimumLatency),
	)
	fmt.Println(
		"Max latency:",
		formatDuration(summary.maximumLatency),
	)
	fmt.Printf(
		"Requests/sec: %.0f\n",
		summary.requestsPerSecond,
	)
}
