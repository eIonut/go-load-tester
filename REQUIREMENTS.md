# Go Load Tester — MVP Requirements

## Goal

Build a simple CLI HTTP load testing tool in Go.

The tool should send multiple HTTP requests concurrently to a target URL and display basic performance statistics.

## Requirements

### 1. CLI Input

- [x] Target URL
- [x] Total number of requests
- [x] Number of concurrent workers
- [x] Validate that requests > 0
- [x] Validate that workers > 0
- [x] Validate that the supplied URL is a valid HTTP/HTTPS URL

Example:

```bash
go-load --url https://example.com --requests 10000 --workers 50
```

### 2. HTTP Requests

- [x] Support `GET` requests only for the MVP
- [x] Send the configured number of requests to the target URL
- [x] Measure the duration of every request
- [x] Use a reusable `http.Client`
- [x] Configure a request timeout

### 3. Concurrency

Use Go concurrency primitives:

- [x] Goroutines
- [x] Worker pool
- [x] `jobs` channel for distributing requests
- [x] `results` channel for collecting results
- [x] Limit the number of concurrent requests using the configured worker count
- [x] Avoid starting more workers than requests

### 4. Request Result

For every request, collect:

- [x] HTTP status code
- [x] Request duration / latency
- [x] Error, if one occurred
- [x] Detect transport/network failures
- [x] Treat unsuccessful HTTP status codes as failures

### 5. Final Statistics

After all requests are completed, display:

- [x] Total requests
- [x] Successful requests
- [x] Failed requests
- [x] Total execution time
- [x] Average latency
- [x] Requests per second
- [x] Minimum latency
- [x] Maximum latency

Example output:

```text
Load test completed

Total requests:     10000
Successful:          9874
Failed:               126

Total duration:      12.40s
Average latency:     58.20ms
Min latency:         31.40ms
Max latency:         412.80ms
Requests/sec:        806.40
```

### 6. Error Handling

Handle basic errors such as:

- [x] Invalid URL format
- [x] Server unavailable / connection errors
- [x] Request timeout
- [x] HTTP request errors
- [x] HTTP error status codes (`4xx`, `5xx`)
- [ ] Prevent division by zero or invalid statistics when no requests succeed

## Remaining MVP Work

- [x] Add proper URL validation
- [x] Create and reuse an `http.Client`
- [x] Add an HTTP request timeout
- [x] Calculate minimum latency
- [x] Calculate maximum latency
- [x] Clean up `LoadTestStatistics` so it represents only one request
- [x] Improve CLI output formatting
- [ ] Test with different combinations of requests/workers
- [ ] Test failure scenarios: invalid URL, unreachable server, timeout, `404`, `500`
- [ ] Run `go vet` on the project
- [ ] Format the project with `gofmt`

## Out of Scope for MVP

Do **not** implement yet:

- POST / PUT / PATCH / DELETE requests
- Custom headers
- Request bodies / JSON
- Authentication
- Latency percentiles (`p50`, `p95`, `p99`)
- Rate limiting / requests-per-second target
- Multiple target URLs
- Charts
- GUI
- Export to files
- Configuration files

## First Milestone

- [x] Send multiple GET requests
- [x] Use a worker pool
- [x] Execute requests concurrently
- [x] Collect results through channels
- [x] Display basic statistics after all requests finish

**First milestone completed.**

## MVP Completion Criteria

The MVP is complete when the program can:

- [x] Receive URL, request count, and worker count from CLI
- [x] Validate the provided configuration
- [x] Execute all configured requests with controlled concurrency
- [x] Handle network errors, HTTP errors, and timeouts without crashing
- [x] Measure individual request latency
- [x] Measure total load test execution time
- [x] Calculate success and failure counts
- [x] Calculate average, minimum, and maximum latency
- [x] Calculate requests per second
- [x] Display a clean final summary
