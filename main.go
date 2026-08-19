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
	totalRequests   int
	successRequests int
	failedRequests  int
	averageLatency  time.Duration
	minLatency      time.Duration
	maxLatency      time.Duration
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

func makeApiRequest(url string, ch chan LoadTestStatistics) {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		ch <- LoadTestStatistics{
			err: err,
		}
		return
	}

	defer resp.Body.Close()

	elapsedTime := time.Since(start)

	ch <- LoadTestStatistics{statusCode: resp.StatusCode, elapsedTime: elapsedTime, err: nil}
}

func main() {
	statistics := make(chan LoadTestStatistics)

	requestsNumber := flag.Int("requests", defaultRequestsNumber, "Total number of requests to be executed")
	targetURL := flag.String("url", "", "Target URL to stress test")
	workersNumber := flag.Int("workers", defaultWorkersNumber, "The number of workers to execute jobs")

	flag.Parse()

	userInput, err := readUserInput(LoadTestArgs{requestsNumber: requestsNumber, targetURL: targetURL, workersNumber: workersNumber})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	results := make([]LoadTestStatistics, 0, userInput.requestsNumber)

	for range userInput.requestsNumber {
		go makeApiRequest(userInput.targetURL, statistics)
	}

	for range userInput.requestsNumber {
		results = append(results, <-statistics)
	}

	//TODO: add a summary
	for _, res := range results {
		fmt.Println(res)
	}

}
