package main

import (
	"flag"
	"fmt"
	"time"
)

const defaultRequestsNumber = 100
const defaultWorkersNumber = 10

func main() {
	requestsNumber := flag.Int(
		"requests",
		defaultRequestsNumber,
		"Total number of requests to be executed",
	)

	targetURL := flag.String(
		"url",
		"",
		"Target URL to stress test",
	)

	workersNumber := flag.Int(
		"workers",
		defaultWorkersNumber,
		"The number of workers to execute jobs",
	)

	flag.Parse()

	userInput, err := readUserInput(LoadTestArgs{
		requestsNumber: requestsNumber,
		targetURL:      targetURL,
		workersNumber:  workersNumber,
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

	requestsPerSecond :=
		float64(summary.totalRequests) /
			summary.totalTestDuration.Seconds()

	fmt.Println("Load test completed")
	fmt.Println()
	fmt.Println("Total requests:", summary.totalRequests)
	fmt.Println("Successful:", summary.successRequests)
	fmt.Println("Failed:", summary.failedRequests)
	fmt.Println()
	fmt.Printf(
		"Total test duration: %.2f ms\n",
		summary.totalTestDuration.Seconds()*1000,
	)
	fmt.Printf(
		"Average latency: %.2f ms\n",
		summary.averageLatency.Seconds()*1000,
	)
	fmt.Printf(
		"Requests/sec: %.2f\n",
		requestsPerSecond,
	)
}
