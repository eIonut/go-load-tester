package main

import (
	"fmt"
	"time"
)

const defaultRequestsNumber = 100
const defaultWorkersNumber = 10

func main() {
	commandFlags := getCommandFlags()

	userInput, err := readUserInput(LoadTestArgs{
		requestsNumber: commandFlags.requestsNumber,
		targetURL:      commandFlags.targetURL,
		workersNumber:  commandFlags.workersNumber,
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
		go worker(jobs, results, userInput.targetURL)
	}

	for i := 0; i < userInput.requestsNumber; i++ {
		statistics = append(statistics, <-results)
	}

	totalTestDuration := time.Since(start)

	summary := aggregateStatistics(statistics)
	summary.totalTestDuration = totalTestDuration
	summary.requestsPerSecond =
		float64(summary.totalRequests) /
			summary.totalTestDuration.Seconds()

	printLoadTestSummary(summary)

}
