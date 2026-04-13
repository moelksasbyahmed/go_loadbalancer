package server

import (
	"fmt"
	"net/url"
	"sync"
	"sync/atomic"

	internal "github.com/moelksasbyahmed/go_loadbalancer/internal"
)

type LoadBalancerUnit struct {
	backend *Backend
	balance serverbalance
}
type LoadBalancer struct {
	ServerPool []*LoadBalancerUnit
	Algorithim BalancerAlgorithm
	mux        sync.RWMutex
	config     *LoadBalancerConfig
}

type LoadBalancerConfig struct {
	Port       string
	Endpoint   string
	Algorithim string
}

type serverbalance struct {
	overalltraffic         atomic.Int64
	current_traffic        atomic.Int64
	max_request_per_server int
}

func NewloadBalancer(algorithm BalancerAlgorithm, config *LoadBalancerConfig) *LoadBalancer {
	return &LoadBalancer{
		ServerPool: make([]*LoadBalancerUnit, 0),
		Algorithim: algorithm,
		config:     config,
	}
}

func (lb *LoadBalancer) AddBackend(backend *Backend, s *serverbalance) error {

	return nil
}
func (lb *LoadBalancer) Populate_LoadBalancer(conf *internal.Config) {

	lb.mux.Lock()
	defer lb.mux.Unlock()
	for _, server := range conf.Servers {
		lb.ServerPool = append(lb.ServerPool, &LoadBalancerUnit{
			backend: &Backend{
				name: server.Name,
				url: func() *url.URL {
					parsedUrl, err := url.Parse(server.URl)
					if err != nil {
						fmt.Printf("Error parsing Url %s /n ", server.URl)
					}
					return parsedUrl
				}(),
				Alive: atomic.Bool{},
			},
			balance: serverbalance{
				overalltraffic:         atomic.Int64{},
				current_traffic:        atomic.Int64{},
				max_request_per_server: server.MaxRequest,
			},
		})
	}

}
