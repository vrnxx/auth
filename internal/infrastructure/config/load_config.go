package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/vrnxx/auth/pkg/env"
)

const DefaultConfigPath = "../configs/app/dev.toml"

func LoadConfig(c interface{}, absolutePath string) {
	var pathConf string
	if absolutePath == "" {
		if envPath := env.GetEnv("CONFIG_PATH", ""); envPath != "" {
			pathConf = envPath
		} else {
			pathConf = DefaultConfigPath
		}
	} else {
		pathConf = absolutePath
	}
	_, err := toml.DecodeFile(pathConf, c)
	if err != nil {
		fmt.Println(color.RedString(err.Error()))
	}
}
