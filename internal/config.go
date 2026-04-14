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
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Endpoint   string `mapstructure:"endpoint"`
	Algorithim string `mapstructure:"algorithim"`
}
type Config struct {
	Servers     []Serversconfig    `mapstructure:"servers"`
	ProxyConfig LoadBalancerConfig `mapstructure:"LoadBalancerConfig"`
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
