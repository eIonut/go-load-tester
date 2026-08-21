package main

import "net/http"

func worker(
	jobs <-chan struct{},
	results chan<- LoadTestStatistics,
	url string,
	method string,
	body string,
	header string,
	client *http.Client,
) {
	for range jobs {
		results <- makeApiRequest(url, method, body, header, client)
	}
}
