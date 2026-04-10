package internal

import (
	"fmt"

	_ "github.com/moelkasabyahmed/go_loadbalancer/internal/server"
	"github.com/spf13/viper"
)

type Serversconfig struct {
	Name        string `mapstructure:"name"`
	URl         string `mapstructure:"url"`
	Alive       bool   `mapstructure:"alive"`
	max_request int    `mapstructure:"max_request_per_server"`
}
type ProxyConfig struct {
	Proxy_port string `mapstructure:"port"`
	Endpoint   string `mapstructure:"endpoint"`
}
type Config struct {
	Servers     []Serversconfig `mapstructure:"servers"`
	ProxyConfig ProxyConfig     `mapstructure:"proxy_server"`
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
