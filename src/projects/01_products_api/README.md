# 01_products_api em container

Este projeto agora pode rodar em container e conversar com o PostgreSQL definido no compose da raiz de `src`.

## Pré-requisitos

- Variáveis no arquivo `.env` em `src`:
  - `POSTGRES_USER`
  - `POSTGRES_PASSWORD`
  - `POSTGRES_DB`
  - `POSTGRES_PORT`
  - `JWT_SECRET` (mínimo 32 caracteres)
  - `JWT_ISSUER` (opcional, padrão: `01_products_api`)
  - `JWT_TTL_MINUTES` (opcional, padrão: `15`)
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

## Autenticação JWT

- `POST /api/v1/auth/register` é rota pública para criar usuário.
- `POST /api/v1/auth/login` é rota pública para autenticar usuário cadastrado.
- Todas as rotas de produto (`/api/v1/products`) exigem token Bearer válido.
- O token usa `HS256`, possui `issuer`, `iat` e `exp`.

### Exemplo de cadastro

```bash
curl -X POST http://localhost:8000/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"senha_forte_123"}'
```

### Exemplo de login

```bash
curl -X POST http://localhost:8000/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"senha_forte_123"}'
```

### Exemplo de acesso em rota protegida

```bash
TOKEN="SEU_TOKEN"
curl http://localhost:8000/api/v1/products \
  -H "Authorization: Bearer $TOKEN"
```
