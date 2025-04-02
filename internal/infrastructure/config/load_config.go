package config

import (
	"github.com/BurntSushi/toml"
	_ "github.com/BurntSushi/toml"
	"github.com/vrnxx/auth/pkg/env"
)

const DefaultConfigPath = "../configs/app/dev.toml"

func LoadConfig(c interface{}, absolutePath string) {
	var pathConf string
	if absolutePath == "" {
		if envPath := env.GetEnv("CONFIG_PATH", ""); envPath != "" {
			pathConf = envPath
		}
		pathConf = DefaultConfigPath
	} else {
		pathConf = absolutePath
	}
	_, err := toml.DecodeFile(pathConf, c)
	if err != nil {
		panic(err)
	}
}
