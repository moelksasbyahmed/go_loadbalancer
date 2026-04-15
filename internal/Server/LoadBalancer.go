package server

import (
	"net/url"
	"sync"
	"sync/atomic"
)

type Backend struct {
	name  string
	url   *url.URL
	Alive atomic.Bool
}

type LoadBalancerUnit struct {
	backend *Backend
	balance serverbalance
}
type LoadBalancer struct {
	ServerPool []*LoadBalancerUnit
	mux        sync.RWMutex
	config     *LoadBalancerConfig
}

type LoadBalancerConfig struct {
	Algorithim BalancerAlgorithm
}

type serverbalance struct {
	overalltraffic  atomic.Int64
	current_traffic atomic.Int64
	Max_request     int
}

func NewloadBalancer(config *LoadBalancerConfig) *LoadBalancer {
	return &LoadBalancer{
		ServerPool: make([]*LoadBalancerUnit, 0),
		config:     config,
	}
}
