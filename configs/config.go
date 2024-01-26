package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName     string `mapstructure:"appName"`
	Environment string `mapstructure:"environment"`
	ApiVersion  string `mapstructure:"apiVersion"`
	Server      Server
	Database    Database
	BaseUrl     string `mapstructure:"baseUrl"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Database struct {
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	DBHost     string `mapstructure:"dbHost"`
	DBPort     string `mapstructure:"dbPort"`
	DBName     string `mapstructure:"dbName"`
	SchemaName string `mapstructure:"schemaName"`
	LogLevel   int    `mapstructure:"logLevel"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
