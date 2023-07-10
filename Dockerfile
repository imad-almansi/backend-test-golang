# Build the application from source
FROM golang:1.20.5-bookworm AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
COPY vendor/ ./vendor/

COPY *.go ./
COPY pkg/ ./pkg/

RUN go build -o /facts ./*.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM ubuntu:jammy AS build-release-stage

WORKDIR /

COPY --from=build-stage /facts /facts

EXPOSE 8080

ENTRYPOINT ["/facts"]