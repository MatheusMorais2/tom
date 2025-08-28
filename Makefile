include .env

generate:
	@echo "Running templ generate..."
	templ generate ./internal/templates -watch

run:
	@echo "Starting the server..."
	go run main.go
watch:
	templ fmt . & templ generate -watch -cmd "go run main.go"
all: generate run

create_migration:
	migrate create -ext=sql -dir=internal/adapters/storage/migrations -seq $(name)

migrate_up: 
	migrate -path=internal/adapters/storage/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)" up

migrate_down: 
	migrate -path=internal/adapters/storage/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)" down 
