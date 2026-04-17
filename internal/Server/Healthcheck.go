package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fatih/color"
)

func (lb *LoadBalancer) RunManualCheck(target string) error {
	if target == "" {
		target = "all"
	}
	if target == "all" {
		var wg sync.WaitGroup
		for _, servers := range lb.ServerPool {
			wg.Add(1)
			go func(u *LoadBalancerUnit) {
				defer wg.Done()
				u.Backend.HealthCheck()
			}(servers)
		}
		wg.Wait()
		log.Println(color.GreenString("Health check completed for all backends"))
		return nil
	}
	for _, servers := range lb.ServerPool {
		if servers.Backend.Name == target || servers.Backend.Url.String() == target {
			servers.Backend.HealthCheck()
			return nil
		}
	}
	return fmt.Errorf("target not found")
}

func (lb *LoadBalancer) GetBackendStatus() map[string]bool {
	status := make(map[string]bool)
	for _, server := range lb.ServerPool {
		status[server.Backend.Name] = server.Backend.Alive.Load()
	}
	return status
}

func (lb *LoadBalancer) StartHealthCheckLoop(context context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		defer ticker.Stop()
		log.Println(color.GreenString("starting the Loadbalancer Health Check on interval %s ", interval))
		for {
			select {
			case <-ticker.C:
				err := lb.RunManualCheck("all")
				if err != nil {
					log.Println(color.RedString("Error occurred while running manual check: %v\n", err))
				}
			case <-context.Done():
				fmt.Println(color.GreenString("Health check loop stopped."))
				return
			}
		}
	}()

}
