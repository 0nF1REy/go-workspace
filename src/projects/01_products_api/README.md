# 01_products_api em container

Este projeto agora pode rodar em container e conversar com o PostgreSQL definido no compose da raiz de `src`.

## Pré-requisitos

- Variáveis no arquivo `.env` em `src`:
  - `POSTGRES_USER`
  - `POSTGRES_PASSWORD`
  - `POSTGRES_DB`
  - `POSTGRES_PORT`
- Engine de container instalada:
  - padrão: `podman`
  - opcional: `docker` (sobrescrevendo `CONTAINER_ENGINE`)

## Build da imagem

A partir de `src`:

make api-image

Isso executa o atalho equivalente a:

podman build -f projects/01_products_api/Dockerfile -t go-workspace-products-api:latest .

Para usar Docker:

make api-image CONTAINER_ENGINE=docker

## Subir API + PostgreSQL

A partir de `src`:

make api-up

- Serviço do banco: `workspace_db`
- Serviço da API: `products_api`
- A API sobe em `http://localhost:8000`
- Teste rápido: `GET /api/v1/ping`

## Logs

make api-logs

## Parar serviços

Somente banco:

make db-down

Derrubar stack completa (API + DB + rede):

make api-down

## Observações de operação

- No compose, a API usa `POSTGRES_HOST=workspace_db` e `POSTGRES_PORT=5432`, garantindo comunicação container a container na rede interna.
- O Dockerfile é multi-stage e usa imagem final distroless nonroot para reduzir superfície de ataque e tamanho.
