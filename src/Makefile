.PHONY: help start build dev prod menu tidy gazelle clean deps-sync db-up db-down

# Default target
help:
	@echo "=== Go Workspace Makefile ==="
	@echo ""
	@echo "Available targets:"
	@echo "  make deps-sync   - Sincroniza dependências Go e Bazel (go mod tidy + bazel mod tidy + gazelle)"
	@echo "  make gazelle     - Executa o gerador de código do Bazel (gazelle)"
	@echo "  make build       - Compila a aplicação"
	@echo "  make dev         - Executa com hot reload (Air + Bazel + DB check)"
	@echo "  make prod        - Executa em modo interativo (menu)"
	@echo "  make start       - Workflow completo: deps-sync, build e prod"
	@echo "  make db-up       - Inicia o banco de dados PostgreSQL"
	@echo "  make db-down     - Para o banco de dados PostgreSQL"
	@echo "  make clean       - Limpa os artefatos de build do Bazel"
	@echo ""
	@echo "Workflow de desenvolvimento:"
	@echo "  1. make db-up"
	@echo "  2. make dev      (hot reload) ou make prod (modo interativo)"
	@echo ""

# Sync all dependencies
deps-sync:
	@echo "=== Sincronizando dependências ==="
	go mod tidy
	bazel mod tidy
	bazel run //:gazelle

# Run gazelle only
gazelle:
	@echo "=== Executando Bazel gazelle ==="
	bazel run //:gazelle

# Build the application
build:
	@echo "=== Compilando aplicação ==="
	bazel build //cmd/main

# Run the application in development mode (hot reload + DB check)
dev:
	@echo "=== Executando modo desenvolvimento (hot reload) ==="
	./dev.sh

# Run the application in production mode (interactive menu)
prod:
	@echo "=== Executando modo produção (interativo) ==="
	./prod.sh

# Run the interactive menu (no hot reload)
menu:
	@echo "=== Executando menu interativo ==="
	bazel run //cmd/main

# Full workflow: deps-sync, build, and prod
start: deps-sync build prod

# Database operations
db-up:
	@echo "=== Iniciando banco de dados PostgreSQL ==="
	podman compose up

db-down:
	@echo "=== Parando banco de dados PostgreSQL ==="
	podman compose down

# Clean build artifacts
clean:
	@echo "=== Limpando artefatos do Bazel ==="
	bazel clean
	rm -rf bazel-*
	rm -rf tmp/
