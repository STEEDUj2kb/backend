.PHONY: clean critic security lint test build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

docker.postgres:
	docker-compose -f docker/compose.yml --env-file=./.env.$(stage) up -d postgres 

compose.build:
	docker-compose -f docker/compose.yml --env-file=./.env.$(stage) build

compose.up: swag
	docker-compose -f docker/compose.yml --env-file=./.env.$(stage) up

compose.down:
	docker-compose -f docker/compose.yml down

swag:
	swag init