# Go Load Tester

A lightweight command-line HTTP load testing tool written in Go.

It sends HTTP requests concurrently using a worker pool and reports request counts, latency, throughput, and aggregated errors.

## Features

- Concurrent HTTP load testing with goroutines
- Worker pool with configurable concurrency
- Configurable total request count
- Supports `GET`, `POST`, `PUT`, `PATCH`, and `DELETE`
- Optional request body
- Optional custom header
- Reusable `http.Client`
- 5-second request timeout
- URL validation
- HTTP `4xx` and `5xx` responses treated as failures
- Network and timeout error handling
- Aggregated error reporting
- ANSI-colored CLI output
- Performance metrics:
  - Successful requests
  - Failed requests
  - Total test duration
  - Average latency
  - Minimum latency
  - Maximum latency
  - Successful requests/second
  - Failed requests/second

## Requirements

- Go 1.22+ recommended

## Run

Run the project without generating a binary:

```bash
go run . --url https://example.com
```

Build an executable:

```bash
go build .
```

## CLI Options

| Flag         | Description                                   | Default           |
| ------------ | --------------------------------------------- | ----------------- |
| `--url`      | Target HTTP/HTTPS URL                         | required          |
| `--requests` | Total number of requests                      | `100`             |
| `--workers`  | Maximum number of concurrent workers          | `10`              |
| `--method`   | HTTP method: GET, POST, PUT, PATCH, DELETE    | `GET` recommended |
| `--body`     | Optional request body                         | empty             |
| `--header`   | Optional custom header in `Key: Value` format | empty             |

## Examples

### GET

```bash
go run . \
  --method GET \
  --url https://jsonplaceholder.typicode.com/posts \
  --requests 1000 \
  --workers 50
```

### POST with JSON body

```bash
go run . \
  --method POST \
  --url https://jsonplaceholder.typicode.com/posts \
  --body '{"title":"hello","body":"load test","userId":1}' \
  --header "Content-Type: application/json" \
  --requests 100 \
  --workers 10
```

### PUT

```bash
go run . \
  --method PUT \
  --url https://jsonplaceholder.typicode.com/posts/1 \
  --body '{"title":"updated"}' \
  --header "Content-Type: application/json" \
  --requests 100 \
  --workers 10
```

### DELETE

```bash
go run . \
  --method DELETE \
  --url https://jsonplaceholder.typicode.com/posts/1 \
  --requests 100 \
  --workers 10
```

## Example Output

```text
Load test completed

Requests:

Total requests:        1000
Successful:             900
Failed:                 100

Latency:

Total test duration:   1.01s
Average latency:       100.77ms
Min latency:            17.90ms
Max latency:           203.43ms

Throughput:

Successful Requests/sec: 891
Failed Requests/sec:      99

Errors:
  100 x unexpected EOF
```

## How It Works

The application creates one job for every configured request and places the jobs in a channel.

A configurable number of worker goroutines consume jobs from the same channel. Each worker executes an HTTP request and sends its result to the results channel.

```text
CLI input
   |
   v
jobs channel
   |
   +----> worker 1 ----+
   +----> worker 2 ----+
   +----> worker 3 ----+--> results channel --> statistics --> summary
   +----> ... ---------+
```

The number of workers controls the maximum concurrency.

For example:

```text
1000 requests + 1 worker
= requests are executed almost sequentially

1000 requests + 100 workers
= up to roughly 100 requests can be in flight concurrently
```

More workers can increase throughput, but can also increase latency and failure rates when the target server reaches its limits.

## Success and Failure Rules

A request is considered successful when:

```text
200 <= status code < 400
```

A request is considered failed when:

- The HTTP client returns a network error
- The request exceeds the configured timeout
- The response status code is `4xx` or `5xx`
- The request cannot be constructed correctly

Errors with the same message are grouped together in the final report.

## Request Timeout

The load tester uses a reusable HTTP client with a 5-second timeout:

```go
&http.Client{
    Timeout: 5 * time.Second,
}
```

If a request exceeds the timeout, it is stopped and counted as a failed request.

## Project Structure

A possible project structure is:

```text
go-load-tester/
├── main.go
├── client.go
├── worker.go
├── input.go
├── stats.go
├── types.go
├── format.go
├── *_test.go
├── go.mod
└── README.md
```

Responsibilities are separated by concern:

- `main.go` — application orchestration
- `client.go` — HTTP request execution, URL validation, HTTP client configuration
- `worker.go` — worker pool logic
- `input.go` — CLI flags and input validation
- `stats.go` — statistics aggregation and summary output
- `types.go` — shared structs
- `format.go` — duration formatting

## Development

Format the project:

```bash
gofmt -w .
```

Run static checks:

```bash
go vet ./...
```

Run tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

## Notes

This tool can generate a significant amount of traffic. Run load or stress tests only against systems you own or have explicit permission to test.

## Possible Future Improvements

The current scope is intentionally small. Possible extensions include:

- Multiple custom headers
- Configurable timeout
- Latency percentiles such as p50, p95, and p99
- Requests-per-second rate limiting
- Authentication
- Export to JSON/CSV
- Configuration files
