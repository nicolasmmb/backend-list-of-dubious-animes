.PHONY: help
# #load env
include .env

build:
	@go build -o ./bin/main ./cmd/server/main.go

run:
	@go run ./cmd/server/main.go

trace:
	@go tool trace ./trace.out 

docker:
	@docker-compose -f ./.infra/docker-compose.yaml --env-file ./.env up -d

migrate-create:
	@migrate create -ext=sql -dir=./.infra/migrations -seq $(name)

migrate-up:
	@migrate -path ./.infra/migrations -database "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up
	# #  force 1
migrate-down:
	@migrate -path ./.infra/migrations -database "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

migrate-force:
	@migrate -path ./.infra/migrations -database "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" force 1