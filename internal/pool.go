package internal

import (
	"sync/atomic"

	server "github.com/moelkasabyahmed/go_loadbalancer/internal/server"
)

type LoadBalancer struct {
	ServerPool map[*server.Backend]*serverbalance
}

type serverbalance struct {
	overalltraffic  atomic.Int64
	max_request     int
	current_traffic atomic.Int64
}

func NewloadBalancer() *LoadBalancer {
	return &LoadBalancer{
		ServerPool: make(map[*server.Backend]*serverbalance),
	}
}

func (lb *LoadBalancer) AddBackend(backend *server.Backend, s *serverbalance) error {

	return nil
}
