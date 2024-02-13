SHELL=/bin/bash
.DEFAULT_GOAL=setup
CURRENTDIR=$(shell dirname `pwd`)

ifneq (,$(wildcard ./.env))
include .env
export
endif

ifneq (,$(wildcard ./.env.test))
include .env.test
export
endif

# Setup application
setup: go.mod
	@echo "`tput bold`#### Verifying configuration files and server certificates ####`tput sgr0`"
	@test -f cert.pem || go run /usr/local/go/src/crypto/tls/generate_cert.go --host localhost
	@test -f .env || cp .env.example .env
	@test -f config.yaml || cp config.example.yaml config.yaml
	@make generate_key
	@echo "## Configuration files are now ready to use ##"

	@sleep 2

	@echo "`tput bold`#### Installing dependencies to your project ####`tput sgr0`"
	go mod tidy

	go get -u golang.org/x/lint/golint
	go install golang.org/x/lint/golint
	go get -u github.com/mgechev/revive@latest
	go install github.com/mgechev/revive@latest

	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
	go install mvdan.cc/gofumpt@latest
	@sleep 2

	@echo "creating .venv and installing it's dependencies"
	test -d .venv || python3 -m venv .venv
	. .venv/bin/activate; pip install pymigratedb

	@echo "## All dependencies installed successfully ##"
	@sleep 2

	@echo ""
	@echo "`tput bold``tput setaf 1`## Verify config.yaml and .env and fill it according to your params ##`tput sgr0`"
	@echo ""

# Run local server
run: .env
	TEST_DATABASE="" go run .

# Generate private.pem file to encrypt in transit data
generate_key:
	test -f private.pem || openssl genpkey -out private.pem -algorithm RSA -pkeyopt rsa_keygen_bits:4096

# Generate public key from private.pem
generate_public:
	test -f private.pem && openssl pkey -in private.pem -pubout -out public.pem

# Run migrations
migrate: .venv
	. .venv/bin/activate; migrate execute --driver pgsql

# Create new migration file in project
create_migration: .venv
	. .venv/bin/activate; migrate create --driver pgsql --migration_name $(name)

update-docker-image:
	docker build -t=ghrc.io/marincor/rinha2024q1:latest .

run-docker-dev:
	docker-compose up -d --build
	docker-compose logs -f

run-docker-prod:
	docker compose up -d

down:
	docker compose down

test:
	./load-test.sh
