package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Server struct {
  Port string `yaml:"port" env-default:"localhost:8090"`
}

type Config struct {
  Env         string  `yaml:"env" env:"ENV" env-default:"develop"`
  StoragePath string  `yaml:"storage-path" env-required:"true"`
  Server              `yaml:"server"`
}

func MustLoad() *Config {
  var configPath string

  configPath = os.Getenv("CONFIG_PATH")

  if configPath == "" {
    flags := flag.String("config", "", "path to the configuration file")
    flag.Parse()

    configPath = *flags

    if configPath == "" {
      log.Fatal("Config path is not set")
    }
  }

  if _, err := os.Stat(configPath); os.IsNotExist(err) {
    log.Fatalf("config file does not exists: %s", configPath)
  }

  var cfg Config

  err := cleanenv.ReadConfig(configPath, &cfg)

  if err != nil {
    log.Fatalf("can not read config file: %s", err.Error())
  }

  return &cfg
}