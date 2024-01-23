.PHONY: help build up down ps %-logs %-exec

REPOSITORY := tus-record
CONTAINERS := backend frontend db

build: ## docker compose build
	docker compose build

create-network:
	docker network create tus-record

up: ## Do docker compose up with hot reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down

ps: ## Check Containers status
	docker compose ps

%-logs: ## Tail docker compose logs $(CONTAINER_NAME)
	docker compose logs $*

%-exec: ## docker exec -it $(CONTAINER_NAME) bash
	docker exec -it $* bash

clean: ## clean docker container and images
	docker rm -f $(CONTAINERS)
	docker rmi -f $(addprefix $(REPOSITORY)-, $(CONTAINERS))

all-clean:
	docker rm -f $$(docker ps -aq)
	docker rmi -f $$(docker images -aq)

init-backend: ## docker create
	$(eval CONTAINER_NAME := $(shell docker create tus-record-backend))
	docker cp $(CONTAINER_NAME):./app/backend/go.mod ./app/backend/
	docker cp $(CONTAINER_NAME):./app/backend/go.sum ./app/backend/
	docker rm -f $(CONTAINER_NAME)

help: ## Show options
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@grep -E '^[a-zA-Z%_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "Examples:"
	@echo "  make backend-logs		# docker compose logs backend"
	@echo "  make frontend-exec		# docker exec -it frontend bash"