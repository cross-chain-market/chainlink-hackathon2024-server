SSH_PRIVATE_KEY=`cat ~/.ssh/id_rsa`
APP_NAME=chainlink-hackathon2024-server
MIGRATE := docker run -v $(shell pwd)/db/migrations:/migrations --network host --rm migrate/migrate:4 -path=migrations/

build: ## Builds Docker image of app
	docker build -t ${APP_NAME}/app -f build/Dockerfile . --build-arg SSH_PRIVATE_KEY="${SSH_PRIVATE_KEY}"

test: ## Run Tests
	go test -p 1 --cover ./...

run-test: up-dev test down ## Run Tests with Docker containers

up: build ## Start all Docker containers
	docker-compose --profile app -p ${APP_NAME} up -d

down: ## Stop Docker containers
	docker-compose -p ${APP_NAME} down

up-dev: ## Start only dev Docker containers
	docker-compose --profile dev -p ${APP_NAME} up -d 

logs: ## Display log output from services
	docker-compose logs -f -t --tail=all

clean: down ## Stop Docker containers and remove volumes 
	docker volume rm ${APP_NAME}_postgres

help:
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

update-deps:
	go get -u ./...
	go mod tidy

migrate-up: ## Run Postgres database up migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) -database postgres://postgres:postgres@localhost:5432/marketplace?sslmode=disable up

migrate-down: ## Down Postgres database migrations
	@echo "Reverting database migrations..."
	@$(MIGRATE) -database postgres://postgres:postgres@localhost:5432/marketplace?sslmode=disable down -all

migrate-drop: ## Drop Postgres database migrations, NOTE don't drop types
	@echo "Dropping database migrations..."
	@$(MIGRATE) -database postgres://postgres:postgres@localhost:5432/marketplace?sslmode=disable drop -f