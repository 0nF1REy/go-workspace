FROM docker.io/library/golang:1.26-alpine AS builder

WORKDIR /app

# Certificates are needed for TLS calls in some environments.
RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/dev ./cmd/dev
COPY projects/01_products_api ./projects/01_products_api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/products-api ./cmd/dev

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app
COPY --from=builder /out/products-api /app/products-api

EXPOSE 8000
ENV GIN_MODE=release

ENTRYPOINT ["/app/products-api"]
