package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type HttpServer struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}
type GRPCServer struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Database struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}
type MarineForecast struct {
	Url string `mapstructure:"url"`
}
type Config struct {
	App            App            `mapstructure:"app"`
	Server         HttpServer     `mapstructure:"http_server"`
	GRPCServer     GRPCServer     `mapstructure:"grpc_server"`
	Database       Database       `mapstructure:"database"`
	MarineForecast MarineForecast `mapstructure:"marine_wave_height_forecast"`
}

func LoadConfig(path string) (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("error when parsing config into struct: %w", err)
	}

	return config, nil
}
