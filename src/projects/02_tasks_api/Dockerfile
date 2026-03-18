FROM --platform=linux/amd64 docker.io/library/golang:1.26-alpine AS builder

RUN apk add --no-cache ca-certificates file

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY internal/ ./internal/
COPY projects/02_tasks_api ./projects/02_tasks_api
COPY projects/02_tasks_api/db/migrations ./db/migrations

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -trimpath -ldflags="-s -w -extldflags '-static'" \
    -o /out/tasks-api ./projects/02_tasks_api/main.go

FROM --platform=linux/amd64 docker.io/library/alpine:latest

RUN addgroup -S nonroot && adduser -S nonroot -G nonroot

WORKDIR /app

COPY --from=builder --chown=nonroot:nonroot /out/tasks-api /app/tasks-api
COPY --from=builder --chown=nonroot:nonroot /app/db/migrations /app/db/migrations

RUN chmod +x /app/tasks-api

USER nonroot

EXPOSE 8000
ENV MIGRATIONS_PATH=/app/db/migrations

ENTRYPOINT ["/app/tasks-api"]
