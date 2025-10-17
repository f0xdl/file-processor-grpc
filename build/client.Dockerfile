# Build app
FROM golang:1.25.1 AS builder
LABEL authors="f0xdl"
WORKDIR /app

ARG CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o build/bin/client ./cmd/client

# Run app
FROM scratch
COPY --from=builder /app/build/bin/client /client

ENTRYPOINT ["/client"]