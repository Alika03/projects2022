package config

import "sync"

var onceConfig sync.Once
var config *Config

type Config struct {
	Server struct {
		Port string `env:"SERVER_PORT"`
	}
	Db struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		DbName   string `env:"DB_NAME"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
	}
	HashParams struct {
		Memory      string `env:"HASH_MEMORY"`
		Iterations  string `env:"HASH_ITERATIONS"`
		Parallelism string `env:"HASH_PARALLELISM"`
		SaltLength  string `env:"HASH_SALT_LENGTH"`
		KeyLength   string `env:"HASH_KEY_LENGTH"`
	}
}

func GetConfig() *Config {
	onceConfig.Do(func() {

	})

	return config
}
