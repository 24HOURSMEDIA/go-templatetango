# Installation

## Install with go

First create a build in the builds directory, then you can run the executable and move it to
/usr/local/bin or include it in your path.

```
go build -ldflags="-s -w" -o builds/%command% .
builds/%command% --help
```

## Install in Docker with a layered Dockerfile

```
FROM %docker_image% AS template_tango

FROM alpine
COPY /%command% /usr/local/bin/%command%
```