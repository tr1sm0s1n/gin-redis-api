# Gin-Redis-API

Gin API for CRUD operations in Redis.

## 🛠 Built With

<div align="left">
<a href="https://go.dev/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/go.svg" width="36" height="36" alt="Go" /></a>
<a href="https://gin-gonic.com/docs/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/gin.svg" width="36" height="36" alt="Gin" /></a>
<a href="https://redis.io/docs/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/redis.svg" width="36" height="36" alt="Redis" /></a>
<a href="https://docs.docker.com/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/docker.svg" width="36" height="36" alt="Docker" /></a>
</div>

## ⚙️ Run Locally

Clone the project

```bash
git clone https://github.com/DEMYSTIF/gin-redis-api.git
cd gin-redis-api
```

Start the database

```bash
docker compose up -d
```

Open [browser](http://localhost:8001) to view the database (see ./config for the credentials)

Start the application

```bash
go run .
```

For live reload, install Air (optional)

```bash
go install github.com/cosmtrek/air@latest
```

Run the application with Air

```bash
air
```
