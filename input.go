package main

import (
	"errors"
	"flag"
)

func readUserInput(input LoadTestArgs) (UserInput, error) {
	if *input.workersNumber <= 0 || *input.requestsNumber <= 0 {
		return UserInput{}, errors.New("you must supply a number of workers and requests > 0")
	}

	if *input.targetURL == "" {
		return UserInput{}, errors.New("no URL supplied")
	}

	return UserInput{
		requestsNumber: *input.requestsNumber,
		targetURL:      *input.targetURL,
		workersNumber:  *input.workersNumber,
	}, nil
}

func getCommandFlags() LoadTestArgs {
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

	return LoadTestArgs{requestsNumber: requestsNumber, targetURL: targetURL, workersNumber: workersNumber}
}
