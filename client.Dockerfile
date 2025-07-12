FROM golang:1.24.4 AS builder
ARG CGO_ENABLED=0
LABEL authors="f0xdl"
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o ./bin/client ./cmd/client/main.go

FROM scratch
COPY --from=builder /src/bin/client /client
ENTRYPOINT ["/client"]