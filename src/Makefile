.PHONY: help start build dev prod menu tidy gazelle clean deps-sync db-up db-down api-image api-up api-down api-logs podman-clean

CONTAINER_ENGINE ?= podman
COMPOSE_CMD ?= $(CONTAINER_ENGINE) compose

# Defina a API pelo comando: make api-up API_NAME=tasks
API_NAME ?= products

# Detecta automaticamente o número do projeto (01, 02...) baseado no nome da API
API_NUMBER := $(shell ls projects | grep -i $(API_NAME)_api | sed 's/_.*//')

# Paths e nomes dinâmicos
API_PROJECT_DIR := projects/$(API_NUMBER)_$(API_NAME)_api
API_IMAGE ?= go-workspace-$(API_NAME)-api:latest
API_SERVICE_NAME ?= $(API_NAME)_api

# ----------------------
# Default target
# ----------------------
help:
	@echo "=== Go Workspace Makefile ==="
	@echo ""
	@echo "Available targets:"
	@echo "  make deps-sync   - Sincroniza dependências Go e Bazel"
	@echo "  make gazelle     - Executa Bazel gazelle"
	@echo "  make build       - Compila a aplicação"
	@echo "  make dev         - Executa API com hot reload"
	@echo "  make prod        - Executa em modo interativo (menu)"
	@echo "  make start       - Workflow completo: deps-sync, build e prod"
	@echo "  make db-up       - Inicia apenas o banco de dados PostgreSQL"
	@echo "  make db-down     - Para o banco de dados PostgreSQL"
	@echo "  make api-image   - Builda a imagem da API (atalho build -t)"
	@echo "  make api-up      - Sobe PostgreSQL + API em containers"
	@echo "  make api-down    - Derruba os containers da stack"
	@echo "  make api-logs    - Acompanha logs da API e do banco"
	@echo "  make clean       - Limpa os artefatos de build do Bazel"
	@echo "  make podman-clean - Para e remove todos os containers podman"
	@echo ""
	@echo "Workflow com containers:"
	@echo "  1. make api-image API_NAME=products"
	@echo "  2. make api-up API_NAME=products"
	@echo "  Exemplo para outra API: API_NAME=tasks"

# ----------------------
# Dependencies
# ----------------------
deps-sync:
	@echo "=== Sincronizando dependências ==="
	go mod tidy
	bazel mod tidy
	bazel run //:gazelle

gazelle:
	@echo "=== Executando Bazel gazelle ==="
	bazel run //:gazelle

# ----------------------
# Build & Run
# ----------------------
build:
	@echo "=== Compilando aplicação ==="
	bazel build //cmd/main

dev:
	@echo "=== Executando modo desenvolvimento (hot reload) ==="
	./dev.sh

prod:
	@echo "=== Executando modo produção (interativo) ==="
	./prod.sh

menu:
	@echo "=== Executando menu interativo ==="
	bazel run //cmd/main

start: deps-sync build prod

# ----------------------
# Database operations
# ----------------------
db-up:
	@echo "=== Iniciando banco de dados PostgreSQL ==="
	$(COMPOSE_CMD) up workspace_db

db-down:
	@echo "=== Parando banco de dados PostgreSQL ==="
	$(COMPOSE_CMD) stop workspace_db

# ----------------------
# Dynamic API container targets
# ----------------------
api-image:
	@echo "=== Buildando imagem da $(API_SERVICE_NAME) (sem cache) ==="
	$(CONTAINER_ENGINE) build --no-cache -f $(API_PROJECT_DIR)/Dockerfile -t $(API_IMAGE) .

api-up:
	@echo "=== Subindo PostgreSQL + $(API_SERVICE_NAME) ==="
	$(COMPOSE_CMD) up workspace_db $(API_SERVICE_NAME)

api-down:
	@echo "=== Parando containers: $(API_SERVICE_NAME) e workspace_db ==="
	$(COMPOSE_CMD) stop $(API_SERVICE_NAME) workspace_db
	@echo "=== Removendo containers específicos ==="
	@# Filtra e remove os containers pelos labels gerados pelo podman-compose
	@for service in $(API_SERVICE_NAME) workspace_db; do \
		ids=$$($(CONTAINER_ENGINE) ps -a -q --filter "label=io.podman.compose.service=$$service"); \
		if [ -n "$$ids" ]; then \
			$(CONTAINER_ENGINE) rm -f $$ids; \
		fi; \
	done

api-logs:
	@echo "=== Logs da $(API_SERVICE_NAME) e PostgreSQL ==="
	$(COMPOSE_CMD) logs -f $(API_SERVICE_NAME) workspace_db

# ----------------------
# Clean build artifacts
# ----------------------
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
