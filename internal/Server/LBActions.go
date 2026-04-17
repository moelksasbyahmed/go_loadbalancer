package server

import (
	"fmt"

	"net/url"
	"slices"

	"sync/atomic"

	internal "github.com/moelksasbyahmed/go_loadbalancer/internal"
	"github.com/spf13/viper"
)

func (lb *LoadBalancer) Populate_LoadBalancer(conf *internal.Config) {

	for i, server := range conf.Servers {
		parsedURL, err := url.Parse(server.URl)
		if err != nil {
			fmt.Printf("Error parsing URL for server %s: %v\n", server.Name, err)
			continue
		}
		lb.AddBackend(&LoadBalancerUnit{
			Backend: &Backend{
				Name:  server.Name,
				Url:   parsedURL,
				Alive: atomic.Bool{},
			},
			Balance: Serverbalance{
				overalltraffic:  atomic.Int64{},
				current_traffic: atomic.Int64{},
				Max_request:     server.MaxRequest,
			},
		})
		lb.ServerPool[i].Backend.Alive.Store(server.Alive) // Set the backend as alive by default  TODO configure a way to
	}

}
func (lb *LoadBalancer) AddBackend(server *LoadBalancerUnit) error {
	lb.mux.Lock()
	defer lb.mux.Unlock()
	for _, s := range lb.ServerPool {
		if s.Backend.Url.String() == server.Backend.Url.String() {
			return fmt.Errorf("backend with url %s already exists", server.Backend.Url)
		}
		if s.Backend.Name == server.Backend.Name {
			return fmt.Errorf("backend with name %s already exists", server.Backend.Name)
		}
		server.Backend.Alive.Store(true)

	}
	lb.ServerPool = append(lb.ServerPool, server)

	return nil
}

func (lb *LoadBalancer) RemoveBackend(server *Backend) error {
	lb.mux.Lock()
	defer lb.mux.Unlock()
	for i, Delserver := range lb.ServerPool {
		if server.Name == Delserver.Backend.Name || server.Url.String() == Delserver.Backend.Url.String() {
			lb.ServerPool = slices.Delete(lb.ServerPool, i, i+1)
			return nil
		}

	}
	return fmt.Errorf("The Server  %s you want to Delete doesnt Exist ", server.Name)
}

func (lb *LoadBalancer) HealthStatus() map[*Backend]bool {
	state := make(map[*Backend]bool)
	for _, server := range lb.ServerPool {
		state[server.Backend] = server.Backend.Alive.Load()
	}
	return state

}

func (lb *LoadBalancer) ViperSync() {
	var currentServers []internal.Serversconfig

	for _, servers := range lb.ServerPool {
		currentServers = append(currentServers, internal.Serversconfig{
			Name:       servers.Backend.Name,
			URl:        servers.Backend.Url.String(),
			Alive:      servers.Backend.IsAlive(),
			MaxRequest: servers.Balance.Max_request,
		})

	}
	viper.Set("servers", currentServers)

}
