build:
	@go build -o bin/GoBase cmd/server/main.go

run: build
	@./bin/GoBase