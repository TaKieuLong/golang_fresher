# name App
APP_NAME = server

run:
	go run ./cmd/${APP_NAME}/

stop:
	docker compose kill

up:
	docker compose up -d

down:
	docker compose down

.PHONY: run

.PHONY: air