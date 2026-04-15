package internal

import (
	"fmt"

	"github.com/spf13/viper"
)

type Serversconfig struct {
	Name       string `mapstructure:"name"`
	URl        string `mapstructure:"url"`
	Alive      bool   `mapstructure:"alive"`
	MaxRequest int    `mapstructure:"max_request_per_server"`
}
type LoadBalancerConfig struct {
	Algorithim string `mapstructure:"algorithim"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
type Config struct {
	Servers            []Serversconfig    `mapstructure:"servers"`
	LoadBalancerConfig LoadBalancerConfig `mapstructure:"LoadBalancerConfig"`
	ServerConfig       ServerConfig       `mapstructure:"ServerConfig"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	fmt.Println("the path is ", path)
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./internal")
	viper.AddConfigPath("..")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var fileconfig Config
	if err := viper.Unmarshal(&fileconfig); err != nil {
		return nil, err

	}
	return &fileconfig, nil

}
