build:
	@go build  -o bin.exe ./cmd/main

run: build
	@.\bin.exe

test:
	@go test -v ./...