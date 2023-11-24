ARG BUILD_IMAGE=golang:1.21-alpine

FROM ${BUILD_IMAGE} AS builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0
#RUN go test ./...
RUN go build -ldflags="-s -w" -o builds/tango .

FROM scratch
COPY --from=builder /app/builds/tango /tango
ENTRYPOINT ["/tango"]