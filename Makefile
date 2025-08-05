GOOSE_DB_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
GOOSE_DIR=sql/schema
DOCKER_COMPOSE=docker-compose -f infra/docker-compose.yml

.PHONY: migrate-up migrate-down migrate-status dc-up dc-down

migrate-up:
	goose -dir $(GOOSE_DIR) postgres "$(GOOSE_DB_URL)" up

migrate-down:
	goose -dir $(GOOSE_DIR) postgres "$(GOOSE_DB_URL)" down

migrate-status:
	goose -dir $(GOOSE_DIR) postgres "$(GOOSE_DB_URL)" status

dc-up:
	$(DOCKER_COMPOSE) up -d

dc-down:
	$(DOCKER_COMPOSE) down