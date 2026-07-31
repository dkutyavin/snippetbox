ifneq (,$(wildcard .env))
include .env
export
endif

.PHONY: run migrate seed create-certs docker-up prepare-env setup check-env

check-env:
	@test -f .env || (echo "Error: .env not found. Run 'make prepare-env' first." >&2 && exit 1)

run: check-env
	go run ./cmd/web
migrate: check-env
	go run ./cmd/migrate
seed: check-env
	go run ./cmd/seed
create-certs:
	mkdir -p tls && cd tls && go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
docker-up:
	docker compose up --wait
prepare-env:
	cp .env.example .env
setup: prepare-env docker-up migrate seed create-certs
