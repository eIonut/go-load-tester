package main

import "time"

type LoadTestArgs struct {
	requestsNumber int
	workersNumber  int
	targetURL      string
}

type UserInput struct {
	requestsNumber int
	workersNumber  int
	targetURL      string
}

type LoadTestStatistics struct {
	statusCode  int
	elapsedTime time.Duration
	err         error
}

type LoadTestSummary struct {
	totalRequests               int
	successRequests             int
	failedRequests              int
	averageLatency              time.Duration
	totalTestDuration           time.Duration
	minimumLatency              time.Duration
	maximumLatency              time.Duration
	failedRequestsPerSecond     float64
	successfulRequestsPerSecond float64
	errorCounts                 map[string]int
}
