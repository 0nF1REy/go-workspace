FROM docker.io/library/golang:1.26-alpine AS builder

WORKDIR /app

# Certificates are needed for TLS calls in some environments.
RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/dev ./cmd/dev
COPY projects/02_tasks_api ./projects/02_tasks_api

COPY projects/02_tasks_api/db/migrations ./db/migrations

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/tasks-api ./cmd/dev

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app
COPY --from=builder /out/tasks-api /app/tasks-api
COPY --from=builder /app/db/migrations /app/db/migrations

EXPOSE 8000
ENV MIGRATIONS_PATH=/app/db/migrations

ENTRYPOINT ["/app/tasks-api"]
