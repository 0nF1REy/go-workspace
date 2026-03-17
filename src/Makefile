.PHONY: help start build dev prod menu tidy gazelle clean deps-sync db-up db-down api-image api-up api-down api-logs podman-clean

CONTAINER_ENGINE ?= podman
COMPOSE_CMD ?= $(CONTAINER_ENGINE) compose
PRODUCTS_API_IMAGE ?= go-workspace-products-api:latest
PRODUCTS_API_DOCKERFILE ?= projects/01_products_api/Dockerfile

# Default target
help:
	@echo "=== Go Workspace Makefile ==="
	@echo ""
	@echo "Available targets:"
	@echo "  make deps-sync   - Sincroniza dependências Go e Bazel (go mod tidy + bazel mod tidy + gazelle)"
	@echo "  make gazelle     - Executa o gerador de código do Bazel (gazelle)"
	@echo "  make build       - Compila a aplicação"
	@echo "  make dev         - Executa API direto com hot reload (Air + Bazel + DB check)"
	@echo "  make prod        - Executa em modo interativo (menu)"
	@echo "  make start       - Workflow completo: deps-sync, build e prod"
	@echo "  make db-up       - Inicia apenas o banco de dados PostgreSQL"
	@echo "  make db-down     - Para o banco de dados PostgreSQL"
	@echo "  make api-image   - Builda a imagem da 01_products_api (atalho para build -t)"
	@echo "  make api-up      - Sobe PostgreSQL + 01_products_api em containers"
	@echo "  make api-down    - Derruba os containers da stack"
	@echo "  make api-logs    - Acompanha logs da API e do banco"
	@echo "  make clean       - Limpa os artefatos de build do Bazel"
	@echo "  make podman-clean - Para e remove todos os containers podman (força e ignora erros)"
	@echo ""
	@echo "Workflow com containers:"
	@echo "  1. make api-image"
	@echo "  2. make api-up"
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
	$(COMPOSE_CMD) up workspace_db

db-down:
	@echo "=== Parando banco de dados PostgreSQL ==="
	$(COMPOSE_CMD) stop workspace_db

# Containers for products API + PostgreSQL
api-image:
	@echo "=== Buildando imagem da 01_products_api ==="
	$(CONTAINER_ENGINE) build -f $(PRODUCTS_API_DOCKERFILE) -t $(PRODUCTS_API_IMAGE) .

api-up:
	@echo "=== Subindo PostgreSQL + 01_products_api ==="
	$(COMPOSE_CMD) up workspace_db products_api

api-down:
	@echo "=== Derrubando stack de containers ==="
	$(COMPOSE_CMD) down

api-logs:
	@echo "=== Logs da 01_products_api e PostgreSQL ==="
	$(COMPOSE_CMD) logs -f products_api workspace_db

# Clean build artifacts
clean:
	@echo "=== Limpando artefatos do Bazel ==="
	bazel clean
	rm -rf bazel-*
	rm -rf tmp/

podman-clean:
	@echo "=== Parando e removendo todos os containers podman ==="
	@echo ""
	@if [ "$$(podman ps -aq)" != "" ]; then \
		podman stop $$(podman ps -aq); \
		podman rm $$(podman ps -aq); \
	else \
		echo "Nenhum container para parar ou remover."; \
		echo ""; \
		podman ps -a; \
	fi
