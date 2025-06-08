codegen-init:
	go run github.com/99designs/gqlgen init

codegen:
	go run github.com/99designs/gqlgen generate

run-dev:
	go run cmd/server/main.go

format:
	goimports -w .

install-goimports:
	go install golang.org/x/tools/cmd/goimports@latest

build:
	go build -o be-pradit-dnd-2025 ./cmd/server 