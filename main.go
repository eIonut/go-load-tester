package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"time"
)

type LoadTestArgs struct {
	requestsNumber *int
	workersNumber  *int
	targetURL      *string
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
	totalRequests     int
	successRequests   int
	failedRequests    int
	averageLatency    time.Duration
	totalTestDuration time.Duration
}

const defaultRequestsNumber = 100
const defaultWorkersNumber = 10

func readUserInput(input LoadTestArgs) (UserInput, error) {

	if *input.workersNumber <= 0 || *input.requestsNumber <= 0 {
		return UserInput{}, errors.New("You must supply a number of workers and requests > 0")
	}

	if *input.targetURL == "" {
		return UserInput{}, errors.New("No URL supplied")
	}

	return UserInput{requestsNumber: *input.requestsNumber, targetURL: *input.targetURL, workersNumber: *input.workersNumber}, nil
}

func makeApiRequest(url string) LoadTestStatistics {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Err:", err)
		return LoadTestStatistics{
			err: err,
		}
	}

	defer resp.Body.Close()

	return LoadTestStatistics{
		statusCode:  resp.StatusCode,
		elapsedTime: time.Since(start),
	}
}

func worker(jobs <-chan struct{}, results chan<- LoadTestStatistics, url string) {
	for range jobs {
		result := makeApiRequest(url)
		results <- result
	}
}

func aggregateStatistics(statistics []LoadTestStatistics) LoadTestSummary {
	summary := LoadTestSummary{totalRequests: len(statistics)}

	var totalLatency time.Duration

	for _, s := range statistics {
		if s.err != nil {
			summary.failedRequests++
		} else {
			summary.successRequests++
			totalLatency += s.elapsedTime
		}
	}

	if summary.successRequests > 0 {
		summary.averageLatency = totalLatency / time.Duration(summary.successRequests)
	}

	return summary
}

func main() {
	requestsNumber := flag.Int("requests", defaultRequestsNumber, "Total number of requests to be executed")
	targetURL := flag.String("url", "", "Target URL to stress test")
	workersNumber := flag.Int("workers", defaultWorkersNumber, "The number of workers to execute jobs")
	flag.Parse()

	userInput, err := readUserInput(LoadTestArgs{requestsNumber: requestsNumber, targetURL: targetURL, workersNumber: workersNumber})
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	jobs := make(chan struct{}, userInput.requestsNumber)
	for range userInput.requestsNumber {
		jobs <- struct{}{}
	}
	close(jobs)

	results := make(chan LoadTestStatistics)
	workers := min(userInput.workersNumber, userInput.requestsNumber)
	statistics := make([]LoadTestStatistics, 0, userInput.requestsNumber)

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
	requestsPerSecond := float64(summary.totalRequests) / summary.totalTestDuration.Seconds()

	fmt.Println("Load test completed")
	fmt.Println("")
	fmt.Println("Total requests:", summary.totalRequests)
	fmt.Println("Successful:", summary.successRequests)
	fmt.Println("Failed:", summary.failedRequests)
	fmt.Println("")
	fmt.Printf("Total test duration: %.2f ms\n", summary.totalTestDuration.Seconds()*1000)
	fmt.Printf("Average latency: %.2f ms\n", summary.averageLatency.Seconds()*1000)
	fmt.Println("Requests/sec:", int(requestsPerSecond))

}
