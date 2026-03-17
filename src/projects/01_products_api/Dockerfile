FROM --platform=linux/amd64 docker.io/library/golang:1.26-alpine AS builder

RUN apk add --no-cache ca-certificates file

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY internal/ ./internal/
COPY projects/01_products_api ./projects/01_products_api
COPY projects/01_products_api/db/migrations ./db/migrations

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -trimpath -ldflags="-s -w -extldflags '-static'" \
    -o /out/products-api ./projects/01_products_api/main.go

RUN file /out/products-api && ls -lh /out/products-api

FROM --platform=linux/amd64 docker.io/library/alpine:latest

RUN addgroup -S nonroot && adduser -S nonroot -G nonroot

WORKDIR /app

COPY --from=builder --chown=nonroot:nonroot /out/products-api /app/products-api
COPY --from=builder --chown=nonroot:nonroot /app/db/migrations /app/db/migrations

RUN chmod +x /app/products-api

USER nonroot

EXPOSE 8000
ENV GIN_MODE=release
ENV MIGRATIONS_PATH=/app/db/migrations

ENTRYPOINT ["/app/products-api"]
