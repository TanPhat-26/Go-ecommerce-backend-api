APP_NAME=go-ecommerce-backend-api
DB_URL?=postgres://postgres:postgres@localhost:5432/go_ecommerce?sslmode=disable
MIGRATIONS_PATH=./migrations

.PHONY: run test tidy fmt build docker-up docker-down docker-ps migrate-up migrate-down
run:
	go run ./cmd/server
test:
	go test ./... -count=1
tidy:
	go mod tidy
fmt:
	gofmt -w $$(go env GOPATH 2>NUL) 2>NUL || gofmt -w ./cmd ./global ./internal ./pkg

build:
	go build -o bin/server ./cmd/server

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-ps:
	docker compose ps

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1