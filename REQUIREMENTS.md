# Go Load Tester — MVP Requirements

## Goal

Build a simple CLI HTTP load testing tool in Go.

The tool should send multiple HTTP requests concurrently to a target URL and display basic performance statistics.

## Requirements

### 1. CLI Input

The application should accept:

- Target URL
- Total number of requests
- Number of concurrent workers

Example:

```bash
go-load --url https://example.com --requests 10000 --workers 50
```

### 2. HTTP Requests

- Support `GET` requests only for the MVP.
- Send the configured number of requests to the target URL.
- Measure the duration of every request.

### 3. Concurrency

Use Go concurrency primitives:

- Goroutines
- Worker pool
- `jobs` channel for distributing requests
- `results` channel for collecting results

### 4. Request Result

For every request, collect:

- HTTP status code
- Request duration / latency
- Success or failure
- Error, if one occurred

### 5. Final Statistics

After all requests are completed, display:

- Total requests
- Successful requests
- Failed requests
- Total execution time
- Average latency
- Requests per second

Example output:

```text
Load test completed

Total requests:     10000
Successful:          9874
Failed:               126

Total duration:      12.4s
Average latency:     58ms
Requests/sec:        806.4
```

### 6. Error Handling

Handle basic errors such as:

- Invalid URL
- Server unavailable
- Request timeout
- HTTP request errors

## Out of Scope for MVP

Do **not** implement yet:

- POST / PUT / DELETE requests
- Custom headers
- Request bodies / JSON
- Authentication
- Latency percentiles (`p95`, `p99`)
- Charts
- GUI
- Export to files
- Configuration files

## First Milestone

Implement a program that:

1. Sends **100 GET requests**
2. Uses a **worker pool**
3. Executes requests concurrently
4. Collects results through channels
5. Displays basic statistics after all requests finish
