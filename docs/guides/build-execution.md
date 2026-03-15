# 🛠️ Build & Execução

Esta seção foi movida do README para facilitar a navegação e detalhar os comandos de build e execução do projeto.

## Comandos principais

- `make deps-sync`: Sincroniza dependências Go e Bazel
- `make build`: Compila a aplicação
- `make dev`: Executa API com hot reload (Air + Bazel + DB check)
- `make prod`: Executa em modo interativo (menu)
- `make start`: Workflow completo: deps-sync, build e prod
- `make db-up`: Inicia o banco PostgreSQL
- `make db-down`: Para o banco
- `make api-image`: Builda imagem da 01_products_api
- `make api-up`: Sobe containers da stack
- `make api-down`: Derruba containers
- `make api-logs`: Acompanha logs
- `make clean`: Limpa artefatos de build

## Sequência mínima para rodar

```sh
go mod tidy
bazel mod tidy
bazel run //:gazelle
bazel build //cmd/main
bazel run //cmd/main
```
