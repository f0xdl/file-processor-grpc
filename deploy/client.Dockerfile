# Build
FROM golang:1.25.1-alpine AS build
LABEL authors="f0xdl"
WORKDIR /app

ARG CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    DEBUG=0 

# Preload debugger 
RUN go install \
    -ldflags "-s -w -extldflags '-static'" \
    github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .



RUN if [ "$DEBUG" = "1" ]; then \
      go build -gcflags "all=-N -l" -o build/bin/client ./cmd/client; \
    else \
      go build -o build/bin/client ./cmd/client; \
    fi

EXPOSE 40000
ENTRYPOINT ["dlv", "exec", "/app/build/bin/client", \
            "--headless", \
            "--listen=:40000", \
            "--api-version=2", \
            "--accept-multiclient" \
]

# Production
FROM alpine:latest  AS production

RUN addgroup -S appuser \
 && adduser -S -G appuser -H -s /sbin/nologin appuser

COPY --from=build --chown=appuser:appuser /app/build/bin/client /app 

USER appuser
ENV HTTP_ADDRESS=":8080" \
    GIN_MODE=release

EXPOSE 8080
CMD ["/app"]
