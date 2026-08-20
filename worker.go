package main

import "net/http"

func worker(
	jobs <-chan struct{},
	results chan<- LoadTestStatistics,
	url string,
	client *http.Client,
) {
	for range jobs {
		results <- makeApiRequest(url, client)
	}
}
