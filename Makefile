generate:
	@echo "Running templ generate..."
	templ generate ./pkg/templates -watch

run:
	@echo "Starting the server..."
	go run main.go
watch:
	templ fmt . & templ generate -watch -cmd "go run main.go"
all: generate run
