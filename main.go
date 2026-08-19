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

const defaultRequestsNumber = 100
const defaultWorkersNumber = 10

func main() {

	requestsNumber := flag.Int("requests", defaultRequestsNumber, "Total number of requests to be executed")
	targetURL := flag.String("url", "", "Target URL to stress test")
	workersNumber := flag.Int("workers", defaultWorkersNumber, "The number of workers to execute jobs")

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

	if *input.targetURL == "" {
		return UserInput{}, errors.New("No URL supplied")
	}

	return UserInput{requestsNumber: *input.requestsNumber, targetURL: *input.targetURL, workersNumber: *input.workersNumber}, nil
}
