package main

import (
	"fmt"
	"net/http"
	"time"
)

func makeApiRequest(url string) LoadTestStatistics {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		return LoadTestStatistics{
			err: err,
		}
	}

	defer resp.Body.Close()

	elapsedTime := time.Since(start)

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return LoadTestStatistics{
			statusCode:  resp.StatusCode,
			elapsedTime: elapsedTime,
		}
	}

	return LoadTestStatistics{
		statusCode:  resp.StatusCode,
		elapsedTime: elapsedTime,
		err:         fmt.Errorf("request failed with status code %d", resp.StatusCode),
	}
}
