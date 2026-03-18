# Despliegue en Hostinger VPS

La ruta recomendada para este proyecto es desplegar la API en un `VPS` de Hostinger usando `Docker Manager` y `Traefik` como reverse proxy. Esa eleccion es una inferencia practica basada en la documentacion oficial de Hostinger para proyectos Docker y APIs.

## Flujo recomendado

1. Crear un `VPS` con la plantilla de Docker habilitada.
2. Desplegar primero el proyecto `Traefik` desde el catalogo de Docker de Hostinger.
3. Apuntar el dominio o subdominio al IP del VPS.
4. Subir este repositorio y desplegar `docker-compose.hostinger.yml`.
5. Configurar las variables del archivo `.env` en el panel o en el compose.

## Variables minimas

- `APP_NAME`
- `APP_ENV=production`
- `APP_PORT=8080`
- `APP_DOMAIN=api.tu-dominio.com`
- `DB_ENABLED`
- `DB_HOST`
- `DB_PORT`
- `DB_NAME`
- `DB_USER`
- `DB_PASSWORD`
- `DB_SSLMODE`

## Que hace `docker-compose.hostinger.yml`

- No publica puertos directamente en el host.
- Se conecta a la red `traefik-proxy`.
- Expone la API a traves de labels de Traefik.
- Deja el puerto interno de la aplicacion en `8080`.

## Pasos sugeridos en Hostinger

1. En `VPS -> Docker Manager`, desplegar Traefik.
2. En `Domains`, crear un registro `A` para el subdominio que apunte al IP del VPS.
3. En `VPS -> Security`, confirmar que el trafico web del VPS esta permitido.
4. Crear el proyecto desde repositorio o subir manualmente el compose.
5. Asignar `APP_DOMAIN` con el dominio real que resolvera hacia Traefik.

## Referencias oficiales

- Docker Manager para VPS:
  https://www.hostinger.com/support/12040789-hostinger-docker-manager-for-vps-simplify-your-container-deployments/
- Traefik con varios proyectos Docker Compose:
  https://www.hostinger.com/support/connecting-multiple-docker-compose-projects-using-traefik-in-hostinger-docker-manager/
- Apuntar dominio al VPS:
  https://www.hostinger.com/support/1583227-how-to-point-a-domain-to-your-vps-at-hostinger/
- Panel de VPS:
  https://www.hostinger.com/support/5726606/
