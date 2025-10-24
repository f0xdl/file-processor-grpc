# Build
FROM golang:1.25.1-alpine AS build
LABEL authors="f0xdl"
WORKDIR /app

ARG CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    DEBUG=0

# Preload Debugger 
RUN go install \
    -ldflags "-s -w -extldflags '-static'" \
    github.com/go-delve/delve/cmd/dlv@latest

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.38 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-${GOOS}-${GOARCH} && \
    chmod +x /bin/grpc_health_probe

COPY go.mod go.sum ./
RUN go mod download
COPY . .



RUN if [ "$DEBUG" = "1" ]; then \
      go build -gcflags "all=-N -l" -o build/bin/file-service ./cmd/file-service; \
    else \
      go build -o build/bin/file-service ./cmd/file-service; \
    fi
    
EXPOSE 40000
ENTRYPOINT ["dlv", "exec", "/app/build/bin/file-service", \
            "--headless", \
            "--listen=:40000", \
            "--api-version=2", \
            "--accept-multiclient" \
]

# Production
FROM alpine:latest  AS production

RUN addgroup -S appuser \
 && adduser -S -G appuser -H -s /sbin/nologin appuser

COPY --from=build --chown=appuser:appuser /app/build/bin/file-service /app 
COPY --from=build --chown=appuser:appuser /bin/grpc_health_probe /bin/grpc_health_probe
RUN mkdir -p /storage && chown -R appuser:appuser /storage

USER appuser
ENV GRPC_ADDRESS=":50051" \
    STORAGE_DIR="/storage"

EXPOSE 50051
CMD ["/app"]

