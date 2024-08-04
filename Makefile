DOCKER_COMPOSE_FILE = docker-compose.yml

SERVICE_NAME = app

build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

logs:
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f $(SERVICE_NAME)

restart:
	docker-compose -f $(DOCKER_COMPOSE_FILE) restart $(SERVICE_NAME)

clean:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down -v --rmi all --remove-orphans

run:
	make build
	make up

rebuild:
	make clean
	make build
	make up

.PHONY: build up down logs restart clean run rebuild