package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func makeApiRequest(url string, client *http.Client) LoadTestStatistics {
	start := time.Now()

	resp, err := client.Get(url)

	if err != nil {
		elapsedTime := time.Since(start)
		if errors.Is(err, context.DeadlineExceeded) ||
			errors.Is(err, os.ErrDeadlineExceeded) {
			return LoadTestStatistics{
				elapsedTime: elapsedTime,
				err:         fmt.Errorf("request timed out: %w", err),
			}
		}
		return LoadTestStatistics{
			elapsedTime: elapsedTime,
			err:         err,
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

func getConfiguredHTTPClient() *http.Client {
	return &http.Client{Timeout: 5 * time.Second}
}
