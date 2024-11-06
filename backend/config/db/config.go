package config

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
}

func LoadConfig() *Config {
    viper.SetConfigFile(".env") // or use viper.AddConfigPath() and viper.SetConfigName()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }

    return &Config{
        DBHost:     viper.GetString("DB_HOST"),
        DBPort:     viper.GetString("DB_PORT"),
        DBUser:     viper.GetString("DB_USER"),
        DBPassword: viper.GetString("DB_PASSWORD"),
        DBName:     viper.GetString("DB_NAME"),
    }
}