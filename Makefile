TARGET = ./bin/mitocho

.PHONY: run clean build fdev bdev fbuild bbuild

fdev:
	@npm run dev --prefix ./frontend

bdev:
	@go run mitocho.go

fbuild:
	@npm run build --prefix ./frontend

bbuild:
	@go build -o $(TARGET) mitocho.go

run: build
	@./$(TARGET)

build: fbuild bbuild
	
clean:
	@rm -f $(TARGET)
