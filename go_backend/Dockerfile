# syntax=docker/dockerfile:1.4
FROM        --platform=$BUILDPLATFORM golang:1.18-alpine AS builder
WORKDIR     /app
COPY        go.mod go.sum ./
RUN         go mod download
COPY        *.go .
RUN         go test
RUN         go build -o go_app_backend main.go


FROM        scratch
COPY        --from=builder /app/go_app_backend /usr/local/bin/go_app_backend
CMD         ["/usr/local/bin/go_app_backend"]