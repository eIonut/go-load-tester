package main

import (
	"errors"
	"flag"
	"fmt"
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

const DEFAULT_REQUESTS_NUMBER = 100
const DEFAULT_WORKERS_NUMBER = 10

func main() {

	requestsNumber := flag.Int("requests", DEFAULT_REQUESTS_NUMBER, "Total number of requests to be executed")
	targetURL := flag.String("url", "", "Target URL to stress test")
	workersNumber := flag.Int("workers", DEFAULT_WORKERS_NUMBER, "The number of workers to execute jobs")

	flag.Parse()

	userInput, err := readUserInput(LoadTestArgs{requestsNumber: requestsNumber, targetURL: targetURL, workersNumber: workersNumber})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(userInput)
}

func readUserInput(input LoadTestArgs) (UserInput, error) {

	if *input.workersNumber <= 0 || *input.requestsNumber <= 0 {
		return UserInput{}, errors.New("You must supply a number of workers and requests > 0")
	}

	if len(*input.targetURL) == 0 {
		return UserInput{}, errors.New("No URL supplied")
	}

	return UserInput{requestsNumber: *input.requestsNumber, targetURL: *input.targetURL, workersNumber: *input.workersNumber}, nil
}
