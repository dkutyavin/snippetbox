include .env
export

.PHONY: run migrate seed create-certs docker-up prepare-env setup

run:
	go run ./cmd/web

migrate:
	go run ./cmd/migrate

seed:
	go run ./cmd/seed

create-certs:
	mkdir -p tls && cd tls && go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost

docker-up:
	docker compose up --wait

prepare-env:
	cp .env.example .env

setup: prepare-env docker-up migrate seed create-certs
