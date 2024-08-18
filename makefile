build:
	@go build -o bin/go-blog cmd/main.go

test:
	@go test -v ./...

run:build
	@./bin/go-blog
