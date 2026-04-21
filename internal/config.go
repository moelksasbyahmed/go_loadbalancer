package internal

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Serversconfig struct {
	Name       string `mapstructure:"name"`
	URl        string `mapstructure:"url"`
	Alive      bool   `mapstructure:"alive"`
	MaxRequest int    `mapstructure:"maxrequest"`
}
type LoadBalancerconfig struct {
	Port                string        `mapstructure:"port"`
	Host                string        `mapstructure:"host"`
	Algorithim          string        `mapstructure:"algorithim"`
	HealthCheckInterval time.Duration `mapstructure:"health_check_interval"`
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

func LoadConfig(path string) (*Config, *viper.Viper, error) {
	mainconfig := viper.New()
	mainconfig.SetConfigName("config")

	mainconfig.SetConfigType("yaml")
	if path != "" {
		fmt.Println("the path is ", path)
		mainconfig.SetConfigFile(path)
	}

	mainconfig.AddConfigPath(".")
	mainconfig.AddConfigPath("./internal")
	mainconfig.AddConfigPath("../../")
	mainconfig.AddConfigPath("..")
	if err := mainconfig.ReadInConfig(); err != nil {
		return nil, nil, err
	}
	var fileconfig Config
	if err := mainconfig.Unmarshal(&fileconfig); err != nil {
		return nil, nil, err

	}
	Serversconfig := viper.New()
	Serversconfig.SetConfigName("servers")
	Serversconfig.SetConfigType("yaml")
	Serversconfig.AddConfigPath(".")
	Serversconfig.AddConfigPath("./internal")
	Serversconfig.AddConfigPath("../../")
	Serversconfig.AddConfigPath("..")
	if err := Serversconfig.ReadInConfig(); err != nil {
		return nil, nil, err

	}
	if err := Serversconfig.UnmarshalKey("servers", &fileconfig.Servers); err != nil {
		return nil, nil, err
	}

	return &fileconfig, Serversconfig, nil

}
