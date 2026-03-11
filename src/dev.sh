#!/bin/bash

# ==========================
# Script para rodar a API com Air
# ==========================

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

if [[ "${DEV_HOT_RELOAD:-0}" == "1" ]]; then
    echo "Rodando com hot reload (Air)..."
    air -c .air.toml
    exit $?
fi

echo "Rodando em modo interativo (menu)..."
bazel run //cmd/main
