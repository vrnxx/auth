package config

type HttpAPIConfig struct {
	Host          string
	Port          int
	BaseURLPrefix string `toml:"base_url_prefix"`
}

type GrpcAPIConfig struct {
	Port int
}
