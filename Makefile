TARGET = mitocho.exe

run: build
	@./$(TARGET)

templ:
	@templ generate --watch --proxy=http://localhost

tailwind:
	@npx tailwindcss -i ./templ/input.css -o ./templ/static/main.css --watch

build:
	@npx tailwindcss -i ./templ/input.css -o ./templ/static/main.css
	@templ generate
	@go build -o $(TARGET) mitocho.go

clean:
	@rm -f $(TARGET)