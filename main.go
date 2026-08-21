package main

import (
	"fmt"
	"time"
)

const defaultRequestsNumber = 100
const defaultWorkersNumber = 10

func main() {
	client := getConfiguredHTTPClient()
	commandFlags := getCommandFlags()

	userInput, err := readUserInput(LoadTestArgs{
		requestsNumber: commandFlags.requestsNumber,
		targetURL:      commandFlags.targetURL,
		workersNumber:  commandFlags.workersNumber,
		body:           commandFlags.body,
		header:         commandFlags.header,
		method:         commandFlags.method,
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jobs := make(chan struct{}, userInput.requestsNumber)

	for range userInput.requestsNumber {
		jobs <- struct{}{}
	}

	close(jobs)

	results := make(chan LoadTestStatistics)

	workers := min(
		userInput.workersNumber,
		userInput.requestsNumber,
	)

	statistics := make(
		[]LoadTestStatistics,
		0,
		userInput.requestsNumber,
	)

	start := time.Now()

	for range workers {
		go worker(jobs, results, userInput.targetURL,
			userInput.method,
			userInput.body,
			userInput.header, client)
	}

	for i := 0; i < userInput.requestsNumber; i++ {
		statistics = append(statistics, <-results)
	}

	totalTestDuration := time.Since(start)

	summary := aggregateStatistics(statistics)

	summary.totalTestDuration = totalTestDuration
	summary.failedRequestsPerSecond =
		float64(summary.failedRequests) /
			summary.totalTestDuration.Seconds()

	summary.successfulRequestsPerSecond =
		float64(summary.successRequests) /
			summary.totalTestDuration.Seconds()

	printLoadTestSummary(summary)

}
