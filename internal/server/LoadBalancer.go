package server

import (
	"net/url"
	"sync"
	"sync/atomic"

	"github.com/spf13/viper"
)

type Backend struct {
	Name  string
	Url   *url.URL
	Alive atomic.Bool
	mux   sync.RWMutex
}

type LoadBalancerUnit struct {
	Backend *Backend
	Balance Serverbalance
}
type LoadBalancer struct {
	ServerPool    []*LoadBalancerUnit
	Algorithim    BalancerAlgorithm
	mux           sync.RWMutex
	Config        *LoadBalancerConfig
	WritingConfig *viper.Viper
}

type LoadBalancerConfig struct {
	Algorithim BalancerAlgorithm
}

type Serverbalance struct {
	Overalltraffic  atomic.Int64
	Current_traffic atomic.Int64
	Max_request     int
}

func NewloadBalancer(config *LoadBalancerConfig, writer *viper.Viper) *LoadBalancer {
	return &LoadBalancer{
		ServerPool:    make([]*LoadBalancerUnit, 0),
		Algorithim:    config.Algorithim,
		Config:        config,
		WritingConfig: writer,
	}
}
