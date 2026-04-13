package internal

/*
import (
	"sync"
	"sync/atomic"

	proxy "github.com/moelksasbyahmed/go_loadbalancer/internal/proxy"
	server "github.com/moelksasbyahmed/go_loadbalancer/internal/server"
)

type LoadBalancerUnit struct {
	backend *server.Backend
	balance serverbalance
}
type LoadBalancer struct {
	ServerPool []*LoadBalancerUnit
	Algorithim proxy.BalancerAlgorithm
	mux        sync.RWMutex
}

type serverbalance struct {
	overalltraffic  atomic.Int64
	current_traffic atomic.Int64
}

func NewloadBalancer() *LoadBalancer {
	return &LoadBalancer{
		ServerPool: make([]*LoadBalancerUnit, 0),
	}
}

func (lb *LoadBalancer) AddBackend(backend *server.Backend, s *serverbalance) error {

	return nil
}
*/
