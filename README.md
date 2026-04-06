# testEnv

Small Go CLI app that prints user info and optionally waits before exit.

## Environment Variables

- `NAME`: display name (default: `arif`)
- `CHARACTER`: favorite character label (default: `octavia`)
- `RUN_TIME`: seconds to wait before exiting (default: `0`)
  - must be an integer between `0` and `3600`

## Run Locally

```bash
go run app.go
```

Example with custom environment:

```bash
NAME="Jane Doe" CHARACTER="SpiderMan" RUN_TIME="3" go run app.go
```

## Run Tests

```bash
go test ./...
```

## Build Binary

```bash
go build app.go
```

## Build and Run Container

```bash
docker build -t testenv:local .
docker run --rm -e NAME="Jane Doe" -e CHARACTER="SpiderMan" -e RUN_TIME="2" testenv:local
```

## CI/CD Notes

- Workflow runs unit tests before Docker build.
- Docker image scanning is performed with Trivy.
- Published images include `latest`, commit-based, and run-based tags.
