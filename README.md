# Gin-Redis-API

Gin API for CRUD operations in Redis.

## 🛠 Built With

[![Go](https://img.shields.io/badge/go-dodgerblue?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Gin](https://img.shields.io/badge/gin-dodgerblue?style=for-the-badge&logo=go&logoColor=white)](https://gin-gonic.com/)
[![Redis](https://img.shields.io/badge/redis-crimson?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/)
[![Docker](https://img.shields.io/badge/docker-navy?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)

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
go install github.com/cosmtrek/air@latest
```

Run the application with Air

```bash
make air
```

## 📜 License

Click [here](./LICENSE.md).

## 🎗️ Contributing

Click [here](./CONTRIBUTING.md).

## ⚖️ Code of Conduct

Click [here](./CODE_OF_CONDUCT.md).
