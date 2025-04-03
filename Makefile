LOCAL_BIN=$(CURDIR)/bin
DC_NEW := docker compose
DC_OLD := docker-compose

DEV_DEFAULT_CONFIG_FILENAME := "dev.toml"
PROD_DEFAULT_CONFIG_FILENAME := "prod.toml"

# ========================Installing deps & tools============================= #

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

# ==========================Generate GRPC===================================== #

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

# =============================Running======================================== #

grpc-serve:
	go run cmd/grpc/main.go

run-dc-local:
	$(DC_NEW) --project-name auth -f docker-compose.local.yaml up -d --build
down-dc-local:
	$(DC_NEW) --project-name auth -f docker-compose.local.yaml down --remove-orphans


# 🚨WARNING🚨: ОБЯЗАТЕЛЬНО ЗАДАЙ ПЕРЕМЕННЫЕ `CONFIG_PATH`
# (если название конфига не prod.toml,то передай его в переменной `CONFIG_FILENAME`)
# EXAMPLE🎓: CONFIG_PATH=/home/alexander/conf/prod-alex.toml CONFIG_FILENAME=prod-alex.toml make run-dc-prod
run-dc-prod:
	$(DC_NEW) --project-name auth -f docker-compose.production.yaml up -d --build
down-dc-prod:
	$(DC_NEW) --project-name auth -f docker-compose.production.yaml down --remove-orphans
