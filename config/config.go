package config

import (
    "log"
    "github.com/spf13/viper"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
}

func LoadConfig() (*Config, error) {
    viper.SetConfigFile(".env") // Assuming you're using an .env file for config

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    config := &Config{
        DBHost:     viper.GetString("DB_HOST"),
        DBPort:     viper.GetString("DB_PORT"),
        DBUser:     viper.GetString("DB_USER"),
        DBPassword: viper.GetString("DB_PASSWORD"),
        DBName:     viper.GetString("DB_NAME"),
    }

    log.Println("Configuration loaded successfully")
    return config, nil
}
