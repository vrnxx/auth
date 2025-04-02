package config

import (
	db "github.com/vrnxx/auth/internal/infrastructure/db/config"
	api "github.com/vrnxx/auth/internal/presentation/api/config"
)

type AppConfig struct {
	Mode string
}

type Config struct {
	AppConfig         `toml:"app"`
	db.DBConfig       `toml:"db"`
	api.HttpAPIConfig `toml:"http_api"`
	api.GrpcAPIConfig `toml:"grpc_api"`
}
