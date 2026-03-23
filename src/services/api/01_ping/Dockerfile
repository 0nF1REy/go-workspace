# Etapa 1: build
FROM golang:1.20-alpine AS builder
WORKDIR /app

# Copia o código
COPY . .

# Build estático sem CGO
RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

# Etapa 2: container mínimo
FROM scratch
COPY --from=builder /app/api /

# Expõe a porta 8000
EXPOSE 8000

# Comando de entrada
ENTRYPOINT ["/api"]
