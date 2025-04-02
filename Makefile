LOCAL_BIN=$(CURDIR)/bin
DC_NEW := docker compose
DC_OLD := docker-compose
PROD_CONFIG_PATH := ./internal/configs/app/prod.toml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path=api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

grpc-serve:
	go run cmd/grpc/main.go

run-dc:
	CONFIG_PATH=$(PROD_CONFIG_PATH) $(DC_NEW) --project-name auth -f docker-compose.yaml up -d --build
down-dc:
	$(DC_NEW) --project-name auth -f docker-compose.yaml down --remove-orphans