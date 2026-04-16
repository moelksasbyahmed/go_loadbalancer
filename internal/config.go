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
type LoadBalancerconfig struct {
	Port       string `mapstructure:"port"`
	Host       string `mapstructure:"host"`
	Algorithim string `mapstructure:"algorithim"`
}

type Config struct {
	Servers            []Serversconfig    `mapstructure:"servers"`
	LoadBalancerConfig LoadBalancerconfig `mapstructure:"LoadBalancerConfig"`
	Adminconfig        AdminConfig        `mapstructure:"adminConfig"`
}
type AdminConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")

	viper.SetConfigType("yaml")
	if path != "" {
		fmt.Println("the path is ", path)
		viper.SetConfigFile(path)
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("./internal")
	viper.AddConfigPath("../../")
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
