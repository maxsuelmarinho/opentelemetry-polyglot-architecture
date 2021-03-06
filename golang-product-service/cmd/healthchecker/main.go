package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("SERVER_PORT", 8080)

	viper.SetConfigType("yml")
	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Println(errors.Wrap(err, "couldn't read config file"))
	}

	viper.AutomaticEnv()
}

func main() {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/health", viper.GetString("SERVER_PORT")))
	if err != nil || resp.StatusCode != http.StatusOK {
		os.Exit(1)
	}

	os.Exit(0)
}
