package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

func parseURL(rawURL string) error {
	parsedURL, err := url.ParseRequestURI(rawURL)

	if err != nil {
		return errors.New("invalid URL")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("URL must use http or https")
	}

	if parsedURL.Host == "" {
		return errors.New("URL must contain a host")
	}

	return nil
}
