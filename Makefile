SHELL=/bin/bash
.DEFAULT_GOAL=setup
CURRENTDIR=$(shell dirname `pwd`)

# Setup application
setup: go.mod
	@echo "`tput bold`#### Installing dependencies to your project ####`tput sgr0`"
	go mod tidy

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
export DB_POOL=60
run:
	HOST=teste
	
	go run .

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
	docker-compose up -d

down:
	docker-compose down

test:
	./load-test.sh
