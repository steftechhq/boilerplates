build & run it
========================

# v1 imperatively from terminal

```sh
docker build -t simple-node .
docker run --rm -p 3000:3000 simple-node
```

# v2 declaratively docker-compose 
```sh
docker compose up --build -d
docker compose down
```

## notes
* build: . replaces docker build -t simple-node .