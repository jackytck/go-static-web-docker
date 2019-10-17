### Overview
A very simple container in go for hosting static local files.

### Usage (with shared volume)
```text
# .env
STATIC_DIR=static
PORT=:3000
LIST_DIR=false
```

```bash
# single command
docker run --env-file .env -v /etc/group:/etc/group:ro -v /etc/passwd:/etc/passwd:ro -u $( id -u $USER ):$( id -g $USER ) -v /mnt/data/static:/root/static --network host --name static-server -d jackytck/go-static-web-docker:v0.0.1
```

```bash
# docker compose
static-server:
  image: jackytck/go-static-web-docker:v0.0.1
  network_mode: "host"
  ports:
    - "8090:8090"
  user: "${UID}:${GID}"
  environment:
    - PORT=:8090
    - LIST_DIR=false
  restart: always
  volumes:
    - /etc/group:/etc/group:ro
    - /etc/passwd:/etc/passwd:ro
    - /home/jacky/static:/root/static

# export UID=$UID
# export GID=$GID
# docker-compose up -d
```
