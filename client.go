package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func makeApiRequest(
	url string,
	method string,
	body string,
	header string,
	client *http.Client,
) LoadTestStatistics {
	start := time.Now()

	req, err := http.NewRequest(
		method,
		url,
		strings.NewReader(body),
	)

	if err != nil {
		return LoadTestStatistics{
			err: err,
		}
	}

	if header != "" {
		parts := strings.SplitN(header, ":", 2)

		if len(parts) != 2 {
			return LoadTestStatistics{
				err: errors.New("invalid header format"),
			}
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)

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
