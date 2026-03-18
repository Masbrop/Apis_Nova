# Apis_Nova

Base inicial para APIs en Go con arquitectura DDD, orientada a despliegue en Hostinger VPS usando Docker.

## Arquitectura

La estructura parte de tres capas principales:

- `internal/domain`: logica de negocio, contratos y servicios de dominio.
- `internal/infrastructure`: conexiones con tecnologias externas, configuracion y bootstrap.
- `internal/handlers`: entrada HTTP y adaptacion de la capa web.

## Estructura

```text
.
|-- cmd/api
|-- docs
|-- internal
|   |-- domain
|   |   `-- status
|   |-- handlers
|   |   `-- http
|   `-- infrastructure
|       |-- bootstrap
|       |-- config
|       |-- database
|       `-- repositories
|-- Dockerfile
|-- docker-compose.yml
|-- docker-compose.dev.yml
`-- docker-compose.hostinger.yml
```

## Caso base implementado

Se dejo un caso funcional de `status/health` para validar que la arquitectura esta cableada correctamente:

- `GET /health`
- `GET /v1/health`

La respuesta indica:

- estado del servicio
- entorno actual
- timestamp
- estado de dependencias

Si `DB_ENABLED=false`, el chequeo usa un repositorio `noop`.
Si `DB_ENABLED=true`, la aplicacion abre una conexion PostgreSQL mediante `pgx` y la usa como dependencia real de salud.

## Variables de entorno

Parte de `.env.example`:

```env
APP_NAME=apis_nova
APP_ENV=development
APP_PORT=8080
SHUTDOWN_TIMEOUT=10s

DB_ENABLED=false
DB_DRIVER=postgres
DB_HOST=postgres
DB_PORT=5432
DB_NAME=apis_nova
DB_USER=postgres
DB_PASSWORD=postgres
DB_SSLMODE=disable

APP_DOMAIN=api.tu-dominio.com
```

## Ejecucion local

Solo API:

```bash
docker compose up --build
```

API + PostgreSQL local:

```bash
docker compose -f docker-compose.yml -f docker-compose.dev.yml up --build
```

## Despliegue en Hostinger

La documentacion revisada apunta a este flujo recomendado:

1. Usar un `VPS` con `Docker Manager`.
2. Desplegar `Traefik` para manejar `80/443` y certificados.
3. Apuntar el dominio o subdominio al VPS.
4. Desplegar este proyecto con `docker-compose.hostinger.yml`.

Resumen de por que:

- Hostinger Docker Manager permite desplegar proyectos Docker Compose desde URL o YAML.
- Hostinger recomienda usar Traefik cuando varios proyectos comparten el mismo VPS y necesitan `80/443`.
- Traefik enruta trafico por dominio y gestiona SSL automaticamente.

Mas detalle en `docs/hostinger.md`.

## Siguiente paso natural

Con esta base ya podemos empezar a modelar el dominio real del negocio:

- entidades y value objects
- casos de uso de aplicacion
- repositorios concretos por tecnologia
- handlers por contexto
- autenticacion, migraciones y observabilidad
