CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

gen-proto-module:
	./scripts/gen_proto.sh ${CURRENT_DIR}

migration-up:
	migrate -path ./migrations/postgres -database 'postgres://postgres:Muhammad@0.0.0.0:5432/books?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://postgres:Muhammad@0.0.0.0:5432/books?sslmode=disable' down

run:
	go run cmd/main.go
