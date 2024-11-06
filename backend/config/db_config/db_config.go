package db_config

import (
	"log"

	"github.com/spf13/viper"
)

type DBConfig struct {
  DBHost     string
  DBPort     string
  DBUser     string
  DBPassword string
  DBName     string
}

func LoadConfig() *DBConfig {
  viper.SetConfigFile("config/db_config/.env")

  if err := viper.ReadInConfig(); err != nil {
    log.Fatalf("Error reading config file: %v", err)
  }

  return &DBConfig{
    DBHost:     viper.GetString("DB_HOST"),
    DBPort:     viper.GetString("DB_PORT"),
    DBUser:     viper.GetString("DB_USER"),
    DBPassword: viper.GetString("DB_PASSWORD"),
    DBName:     viper.GetString("DB_NAME"),
  }
}