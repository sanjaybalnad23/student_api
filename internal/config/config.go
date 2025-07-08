package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}

type Config struct {
	Env                      string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	PostgresConnectionString string `yaml:"postgres_connection_string" env:"POSTGRES_CONNECTION_STRING"`
	HTTPServer               `yaml:"http_server" env:"HTTP_SERVER"`
}

func LoadConfiguration() *Config {
	var configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists %s", configPath)
	}

	var cfg Config
	var err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file:%s", err.Error())
	}

	return &cfg

}
