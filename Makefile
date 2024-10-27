target=authpod


compile:
	go build -o ./bin/$(target) ./cmd/$(target)/*.go

all: compile

build.docker:
	@docker build -t droptableusers:$(target) -f build/Dockerfile .
	@echo "\nnamed docker.io/library/droptableusers:$(target)\n"

docker.deploy.postgres:
	@docker-compose -f ./deploy/postgres.docker-compose.yml -p droptableusers up -d

postgres.up:
	@GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING=postgres://postgres:postgres@127.0.0.1:5432/postgres \
	goose -dir=./etc/db/migrations/postgres up

postgres.down:
	@GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING=postgres://postgres:postgres@127.0.0.1:5432/postgres \
	goose -dir=./etc/db/migrations/postgres down