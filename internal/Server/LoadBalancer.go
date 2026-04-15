package server

import (
	"net/url"
	"sync"
	"sync/atomic"
)

type Backend struct {
	Name  string
	Url   *url.URL
	Alive atomic.Bool
}

type LoadBalancerUnit struct {
	Backend *Backend
	Balance Serverbalance
}
type LoadBalancer struct {
	ServerPool []*LoadBalancerUnit
	Algorithim BalancerAlgorithm
	mux        sync.RWMutex
	config     *LoadBalancerConfig
}

type LoadBalancerConfig struct {
	Algorithim BalancerAlgorithm
}

type Serverbalance struct {
	overalltraffic  atomic.Int64
	current_traffic atomic.Int64
	Max_request     int
}

func NewloadBalancer(config *LoadBalancerConfig) *LoadBalancer {
	return &LoadBalancer{
		ServerPool: make([]*LoadBalancerUnit, 0),
		Algorithim: config.Algorithim,
		config:     config,
	}
}
