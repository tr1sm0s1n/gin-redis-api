# Gin-Redis-API

Gin API for CRUD operations in Redis.

## 🛠 Built With

[![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=for-the-badge)](https://go.dev/)
[![Gin Badge](https://img.shields.io/badge/Gin-008ECF?logo=gin&logoColor=fff&style=for-the-badge)](https://gin-gonic.com/)
[![Redis Badge](https://img.shields.io/badge/Redis-DC382D?logo=redis&logoColor=fff&style=for-the-badge)](https://redis.io/)
[![Docker Badge](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=fff&style=for-the-badge)](https://www.docker.com/)

## ⚙️ Run Locally

Clone the project

```bash
git clone https://github.com/tr1sm0s1n/gin-redis-api.git
cd gin-redis-api
```

Start the database

```bash
make up
```

Open [browser](http://localhost:8001) to view the database (see ./config for the credentials)

Start the application

```bash
make run
```

For live reload, install Air (optional)

```bash
make air
```

Run the application with Air

```bash
make dev
```
