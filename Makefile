include .env
export

export PROJECT_ROOT=$(shell pwd)

init:
	@mkdir -p $(PROJECT_ROOT)/out/pgdata

env-up: init
	@docker compose up postgres redis kafka ui -d 

env-down:
	@docker compose down postgres redis kafka ui


env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность потери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down postgres redis kafka ui && \
		sudo rm -rf out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder
	
env-port-close:
	@docker compose down  port-forwarder


migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример make migrate-create seq=init"; \
		exit 1; \
	fi;
	docker compose --profile tools run --rm postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример make migrate-action action=down"; \
		exit 1; \
	fi;
	@docker compose --profile tools run --rm postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

services-up:
	@docker-compose up api-gateway notification-worker history-service -d

services-down:
	@docker-compose down api-gateway notification-worker history-service


services-rebuild:
	@docker-compose build --no-cache api-gateway notification-worker history-service

swagger-gen:
	@docker compose --profile tools run --rm swagger \
		init \
		-g api-gateway/cmd/main.go \
		-o api-gateway/docs \
		--parseInternal \
		--parseDependency

ps:
	@docker compose ps


