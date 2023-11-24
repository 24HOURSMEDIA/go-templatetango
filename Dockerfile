FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go test ./...
RUN go build -ldflags="-s -w" -o builds/tango .

FROM scratch
COPY --from=builder /app/builds/tango /tango
ENTRYPOINT ["/tango"]