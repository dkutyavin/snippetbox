include .env
export

.PHONY: run migrate seed

run:
	go run ./cmd/web

migrate:
	go run ./cmd/migrate

seed:
	go run ./cmd/seed
