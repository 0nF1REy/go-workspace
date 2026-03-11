#!/bin/bash

# ==========================
# Production Mode: Interactive Menu (No Hot Reload)
# ==========================

GO_BIN="$HOME/go/bin"

# Adiciona Go bin no PATH se ainda não estiver
case ":$PATH:" in
    *":$GO_BIN:"*) ;;
    *) export PATH="$PATH:$GO_BIN" ;;
esac

echo "Rodando em modo interativo (menu)..."
bazel run //cmd/main
