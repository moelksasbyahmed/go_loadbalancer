package server

import (
	"fmt"
	"net/url"
	"slices"

	"sync/atomic"

	internal "github.com/moelksasbyahmed/go_loadbalancer/internal"
)

func (lb *LoadBalancer) Populate_LoadBalancer(conf *internal.Config) {

	for i, server := range conf.Servers {
		parsedURL, err := url.Parse(server.URl)
		if err != nil {
			fmt.Printf("Error parsing URL for server %s: %v\n", server.Name, err)
			continue
		}
		lb.AddBackend(&LoadBalancerUnit{
			backend: &Backend{
				name:  server.Name,
				url:   parsedURL,
				Alive: atomic.Bool{},
			},
			balance: serverbalance{
				overalltraffic:  atomic.Int64{},
				current_traffic: atomic.Int64{},
				Max_request:     server.MaxRequest,
			},
		})
		lb.ServerPool[i].backend.Alive.Store(server.Alive)
	}

}
func (lb *LoadBalancer) AddBackend(server *LoadBalancerUnit) error {
	lb.mux.Lock()
	defer lb.mux.Unlock()
	for _, s := range lb.ServerPool {
		if s.backend.url.String() == server.backend.url.String() {
			return fmt.Errorf("backend with url %s already exists", server.backend.url)
		}
		if s.backend.name == server.backend.name {
			return fmt.Errorf("backend with name %s already exists", server.backend.name)
		}

	}
	lb.ServerPool = append(lb.ServerPool, server)

	return nil
}

func (lb *LoadBalancer) RemoveBackend(server *LoadBalancerUnit) error {
	lb.mux.Lock()
	defer lb.mux.Unlock()
	for i, Delserver := range lb.ServerPool {
		if server.backend.name == Delserver.backend.name {
			lb.ServerPool = slices.Delete(lb.ServerPool, i, i+1)
			return nil
		}

	}
	return fmt.Errorf("The Server  %s you want to Delete doesnt Exist ", server.backend.name)

}
