package main

import (
	"errors"
	"flag"
)

func readUserInput(input LoadTestArgs) (UserInput, error) {
	if input.workersNumber <= 0 || input.requestsNumber <= 0 {
		return UserInput{}, errors.New("you must supply a number of workers and requests > 0")
	}

	if input.targetURL == "" {
		return UserInput{}, errors.New("no URL supplied")
	}

	if err := parseURL(input.targetURL); err != nil {
		return UserInput{}, err
	}

	return UserInput{
		requestsNumber: input.requestsNumber,
		targetURL:      input.targetURL,
		workersNumber:  input.workersNumber,
		method:         input.method,
		body:           input.body,
		header:         input.header,
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

	method := flag.String(
		"method",
		"GET",
		"The method of the request: GET, POST, PUT, PATCH, DELETE",
	)

	body := flag.String(
		"body",
		"",
		"The body of the request",
	)

	header := flag.String(
		"header",
		"",
		"The header of the request",
	)

	flag.Parse()

	return LoadTestArgs{requestsNumber: *requestsNumber, targetURL: *targetURL, workersNumber: *workersNumber, body: *body, method: *method, header: *header}
}
