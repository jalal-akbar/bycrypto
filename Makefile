build:
	@go build -o ./bin/bycrypt

run:build
	@./bin/docker

test:
	@go test -v ./...