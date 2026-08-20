package main

func worker(
	jobs <-chan struct{},
	results chan<- LoadTestStatistics,
	url string,
) {
	for range jobs {
		results <- makeApiRequest(url)
	}
}
