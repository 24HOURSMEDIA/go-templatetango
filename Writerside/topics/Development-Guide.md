# Development Guide

## Build and Run

### Build and run in docker

Building in docker stores the app file %command% in the root directory of a docker scratch image.
The application has no dependencies, so it can be copied directly from the docker image
in a multi-layered DockerFile.
The docker build also runs all tests, so they should not fail.

```
docker build -t 24hoursmedia/template-tango .
```

```
docker run --rm 24hoursmedia/template-tango --help
docker run --rm 24hoursmedia/template-tango stick:list-filters
```

### Build and run with go

```
go build -ldflags="-s -w" -o builds/%command% .
```

Run with go:

```
go run main.go --help
```

## Run tests

```
go test ./...
```
