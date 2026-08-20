package main

import "errors"

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
