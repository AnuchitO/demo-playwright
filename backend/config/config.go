package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	Server   server
	Database Database
}

type server struct { // TODO: private type
	Host string `env:"BACKEND_SERVER_HOST"`
	Port string `env:"BACKEND_SERVER_PORT" envDefault:"8080"`
}

type Database struct { // TODO: private type
	PostgresURI string `env:"POSTGRES_URI,required"`
}

func Env(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("Environment variable " + key + " is not set")
	}

	return value
}

func ToBoolean(strVal string) bool {
	boolVal, _ := strconv.ParseBool(strVal)
	return boolVal
}

var once sync.Once
var config Config

func C(envPrefix string) Config {
	if len(envPrefix) > 1 {
		log.Fatal("can pass only one prefix for env but your prefix:", envPrefix)
	}

	once.Do(func() {
		var prefix string
		if envPrefix != "" {
			prefix = fmt.Sprintf("%s_", envPrefix)
		}

		opts := env.Options{
			Prefix: prefix,
		}

		db := Database{}
		if err := env.ParseWithOptions(&db, opts); err != nil {
			log.Fatal(err)
		}

		srv := server{}
		if err := env.ParseWithOptions(&srv, opts); err != nil {
			log.Fatal(err)
		}

		h, _ := os.Hostname()
		port := srv.Port
		if port == "" {
			port = os.Getenv("PORT")
		}

		config = Config{
			Server: server{
				Host: h,
				Port: port,
			},
			Database: db,
		}
	})

	return config
}
