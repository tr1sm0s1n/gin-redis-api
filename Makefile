build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

compile:
	go build -o bin/gin-redis-api

start:
	./gin-redis-api

run:
	go run .

air:
	air
