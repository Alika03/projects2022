package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

var onceConfig sync.Once
var config *Config

type Config struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	Db struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		DbName   string `json:"name"`
		User     string `json:"username"`
		Password string `json:"password"`
	} `json:"db"`
	HashParams struct {
		Memory      string `json:"memory"`
		Iterations  string `json:"iterations"`
		Parallelism string `json:"parallelism"`
		SaltLength  string `json:"salt_length"`
		KeyLength   string `json:"key_length"`
	} `json:"hash_params"`
}

func GetConfig() *Config {
	onceConfig.Do(func() {
		dataBytes, err := ioutil.ReadFile("../../tsconfig.json")
		if err != nil {
			panic(err)
		}
		config = &Config{}

		if err = json.Unmarshal(dataBytes, config); err != nil {
			panic(err)
		}
	})

	return config
}
