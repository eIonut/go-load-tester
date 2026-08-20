package main

import "time"

func aggregateStatistics(statistics []LoadTestStatistics) LoadTestSummary {
	summary := LoadTestSummary{
		totalRequests: len(statistics),
	}

	var totalLatency time.Duration

	for _, s := range statistics {
		if s.err != nil {
			summary.failedRequests++
			continue
		}

		summary.successRequests++
		totalLatency += s.elapsedTime
	}

	if summary.successRequests > 0 {
		summary.averageLatency =
			totalLatency / time.Duration(summary.successRequests)
	}

	return summary
}
