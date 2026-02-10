GOOSE_DBSTRING = "root:123456@tcp(127.0.0.1:33306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema 

# name App
APP_NAME = server

dev:
	go run ./cmd/server/main.go

run:
	docker compose up -d && go run  ./cmd/${APP_NAME}

kill:
	docker compose kill

up:
	docker compose up -d

down:
	docker compose down
upse:
	powershell -Command "$$env:GOOSE_DRIVER='mysql'; $$env:GOOSE_DBSTRING='$(GOOSE_DBSTRING)'; goose -dir $(GOOSE_MIGRATION_DIR) up"
downse:
	powershell -Command "$$env:GOOSE_DRIVER='mysql'; $$env:GOOSE_DBSTRING='$(GOOSE_DBSTRING)'; goose -dir $(GOOSE_MIGRATION_DIR) down"
resetse:
	powershell -Command "$$env:GOOSE_DRIVER='mysql'; $$env:GOOSE_DBSTRING='$(GOOSE_DBSTRING)'; goose -dir $(GOOSE_MIGRATION_DIR) reset"

.PHONY : run downse upse resetse

.PHONY: run

.PHONY: air