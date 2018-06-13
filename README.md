### Overview
A very simple container in go for hosting static local files.

### Usage (with shared volume)
``` text
# .env
STATIC_DIR=static
PORT=3000
```

``` bash
docker run --env-file .env -v /etc/group:/etc/group:ro -v /etc/passwd:/etc/passwd:ro -u $( id -u $USER ):$( id -g $USER ) -v /mnt/data/static:/root/static --network host --name static-server -d jackytck/go-static-web-docker:v0.0.1
```
