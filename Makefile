TARGET = ./bin/mitocho

.PHONY: run clean templ tailwind build git-push

dev:
	@go run mitocho.go

git-push:
	@templ generate

run: build
	@./$(TARGET)

templ:
	@templ generate --watch

tailwind:
	@npx tailwindcss -i ./templ/input.css -o ./templ/static/main.css --watch

build:
	@npx tailwindcss -i ./templ/input.css -o ./templ/static/main.css
	@templ fmt .
	@templ generate
	@go build -o $(TARGET) mitocho.go

clean:
	@rm -f $(TARGET)
