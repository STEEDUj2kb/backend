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

docker.run: swag docker.fiber

docker.fiber.build:
	docker build -f docker/go.Dockerfile -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name steedu-fiber \
		-p 5000:5000 \
		fiber

docker.stop: docker.stop.fiber

docker.stop.fiber:
	docker stop steedu-fiber

swag:
	swag init