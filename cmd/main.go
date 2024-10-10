package main

import (
	"github.com/spf13/viper"
)

func main() {
	err := setupConfig("/config/.env")
	if err != nil {
		panic(err)
	}
}

func setupConfig(path string) error {
	viper.SetConfigName(path)
	return viper.ReadInConfig()
}
