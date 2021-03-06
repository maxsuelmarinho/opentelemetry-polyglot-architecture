package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api"
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
	api.StartServer()
}
