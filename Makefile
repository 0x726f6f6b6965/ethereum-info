PROJECTNAME := $(shell basename "$(PWD)")
include .env
export $(shell sed 's/=.*//' .env)

# Protobuf

# proto-lint - linting for Protobuf files 
.PHONY: proto-lint
proto-lint:
	@buf lint

# proto-gen - generate Protobuf related Go files
.PHONY: proto-gen
proto-gen:
	@buf generate

# proto-check-breaking - check breaking
.PHONY: proto-check-breaking
proto-check-breaking:
	@buf breaking --against '.git#branch=main' --error-format=json | jq .

# proto-all - run proto-gen, proto-lint, proto-check-breaking
.PHONY: proto-all
proto-all: proto-gen proto-lint proto-check-breaking

# proto-clean - clean Protobuf related Go files
.PHONY: proto-clean
proto-clean: 
	@find protos -type f -name "*.go" -delete

## DB

.PHONY: set-psql
set-psql:
	@docker run --name MyPostgres -d -p 5432:5432 \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		--rm postgres:latest

.PHONY: set-redis
set-redis:
	@docker run --name redis-lab -p 6379:6379 -d redis
