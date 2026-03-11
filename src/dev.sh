#!/bin/bash

# ==========================
# Development Mode: Hot Reload + DB Check
# ==========================

set -euo pipefail

GO_BIN="$HOME/go/bin"

# Adiciona Go bin no PATH se ainda não estiver
case ":$PATH:" in
    *":$GO_BIN:"*) ;;
    *) export PATH="$PATH:$GO_BIN" ;;
esac

# Verifica se Air está instalado
if ! command -v air &> /dev/null; then
    echo "Air não encontrado. Instalando..."
    go install github.com/air-verse/air@latest || { echo "Erro ao instalar Air"; exit 1; }
    echo "Air instalado com sucesso!"
fi

# Detecta comando de compose (podman ou docker)
compose_cmd=()
container_runtime=""
if command -v podman &> /dev/null; then
    compose_cmd=(podman compose)
    container_runtime="podman"
elif command -v docker &> /dev/null; then
    compose_cmd=(docker compose)
    container_runtime="docker"
fi

# Resolve porta do host a partir do .env quando existir
DB_PORT=5432
if [[ -f .env ]]; then
    env_port=$(grep -E '^POSTGRES_PORT=' .env | tail -n 1 | cut -d'=' -f2 | tr -d '"' || true)
    if [[ -n "${env_port:-}" ]]; then
        DB_PORT="$env_port"
    fi
fi

echo "Verificando banco de dados PostgreSQL..."

db_ok=0

# 1) Checagem via porta do host
if command -v nc &> /dev/null; then
    if nc -z localhost "$DB_PORT" &> /dev/null; then
        db_ok=1
    fi
else
    # Fallback sem netcat
    if (echo > /dev/tcp/127.0.0.1/"$DB_PORT") >/dev/null 2>&1; then
        db_ok=1
    fi
fi

# 2) Checagem via runtime de container (mais confiavel que filtros do compose)
if [[ "$db_ok" -eq 0 ]] && [[ -n "$container_runtime" ]]; then
    container_name="workspace_postgres"

    if "$container_runtime" ps --format '{{.Names}}' 2>/dev/null | grep -qx "$container_name"; then
        # Tenta readiness dentro do container
        if "$container_runtime" exec "$container_name" pg_isready -h 127.0.0.1 -p 5432 >/dev/null 2>&1; then
            db_ok=1
        else
            # Se nao conseguir pg_isready, considera ao menos container up com porta publicada
            if "$container_runtime" ps --format '{{.Names}} {{.Ports}}' 2>/dev/null | grep -q "${container_name} .*:${DB_PORT}->5432/tcp"; then
                db_ok=1
            fi
        fi
    fi
fi

# 3) Fallback via compose (caso nome do container seja alterado)
if [[ "$db_ok" -eq 0 ]] && [[ "${#compose_cmd[@]}" -gt 0 ]]; then
    if "${compose_cmd[@]}" ps 2>/dev/null | grep -Eq 'workspace_db|workspace_postgres'; then
        db_ok=1
    fi
fi

if [[ "$db_ok" -eq 0 ]]; then
    echo ""
    echo "Erro: PostgreSQL nao esta acessivel (porta do host: ${DB_PORT} ou container workspace_db)."
    echo ""
    echo "Certas funcionalidades (especialmente as APIs) nao irao funcionar sem o banco de dados."
    echo "Por favor, execute novamente com o banco ligado:"
    echo ""
    echo "  make db-up"
    echo ""
    echo "Depois rode novamente:"
    echo ""
    echo "  make dev"
    echo ""
    exit 1
fi

echo "PostgreSQL esta ativo"
echo ""
echo "Rodando com hot reload (Air)..."
echo ""

air -c .air.toml
