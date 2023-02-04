# Gouserman

Standalone, simple user management console written in Go , Vue 3, and Postgres.

### Prerequisites

1. Create new file called `jwt.secret` in root folder of docker-compose.yml. This is to save jwt secret upon first generation.

2. Set `POSTGRES_*` in docker-compose.yml to set postgres database details. This defaults to the `run` mode.

3. Set `ADMIN_MASTER_PASSWORD` in docker-compose.yml

### Running

```sh
# First time
cd frontend && yarn install && cd ..

# Then run
./manage.py run
```
