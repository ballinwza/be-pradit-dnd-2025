codegen-init:
	go run github.com/99designs/gqlgen init

codegen:
	go run github.com/99designs/gqlgen generate

run-dev:
	go run cmd/server/main.go