# 🏢 Projetos

Esta seção foi movida do README para facilitar a navegação e detalhar os projetos presentes no workspace.

## 01_products_api

- API REST para produtos, JWT, PostgreSQL
- Pode rodar em container (docker/podman)
- Endpoints: `/api/v1/ping`, `/auth/register`, etc.
- Configuração via `.env` e `docker-compose.yml`
- Build: `make api-image`
- Subir stack: `make api-up`

## ping-api

- Serviço simples de ping HTTP
- Porta padrão: 8080
