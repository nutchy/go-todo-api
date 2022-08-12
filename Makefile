
APP_NAME := go-todo-api

build:
	@docker build -t $(APP_NAME) -f docker/Dockerfile .

start:
	@docker compose -f docker/docker-compose.yaml up -d

stop:
	@docker compose -f docker/docker-compose.yaml down

test:
	curl http://localhost:8888/v1/hello