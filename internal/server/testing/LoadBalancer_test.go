package LB_testing

import (
	"net/url"
	"testing"

	"github.com/moelksasbyahmed/go_loadbalancer/internal/server"
	"github.com/spf13/viper"
)

func TestNewLoadBalancer(t *testing.T) {
	config := &server.LoadBalancerConfig{
		Algorithim: nil,
	}
	vp := viper.New()
	lb := server.NewloadBalancer(config, vp)

	if lb == nil {
		t.Error("LoadBalancer creation failed")
	}
	if len(lb.ServerPool) != 0 {
		t.Errorf("Expected empty ServerPool, got %d", len(lb.ServerPool))
	}
	if lb.Config != config {
		t.Error("Config not properly assigned")
	}
	if lb.WritingConfig != vp {
		t.Error("Viper config not properly assigned")
	}
}

func TestAddBackend(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	url1, _ := url.Parse("http://localhost:8001")
	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
		Balance: server.Serverbalance{
			Max_request: 100,
		},
	}

	err := lb.AddBackend(backend1)
	if err != nil {
		t.Errorf("Failed to add backend: %v", err)
	}
	if len(lb.ServerPool) != 1 {
		t.Errorf("Expected 1 backend, got %d", len(lb.ServerPool))
	}
	if lb.ServerPool[0].Backend.Name != "Server1" {
		t.Errorf("Expected backend name 'Server1', got %s", lb.ServerPool[0].Backend.Name)
	}
}

func TestAddBackendDuplicateURL(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	url1, _ := url.Parse("http://localhost:8001")
	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
	}
	backend2 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server2",
			Url:  url1,
		},
	}

	lb.AddBackend(backend1)
	err := lb.AddBackend(backend2)

	if err == nil {
		t.Error("Expected error for duplicate URL, got nil")
	}
}

func TestAddBackendDuplicateName(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	url1, _ := url.Parse("http://localhost:8001")
	url2, _ := url.Parse("http://localhost:8002")
	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
	}
	backend2 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url2,
		},
	}

	lb.AddBackend(backend1)
	err := lb.AddBackend(backend2)

	if err == nil {
		t.Error("Expected error for duplicate name, got nil")
	}
}

func TestRemoveBackend(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	url1, _ := url.Parse("http://localhost:8001")
	backend := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
	}
	lb.AddBackend(backend)

	err := lb.RemoveBackend(backend.Backend)
	if err != nil {
		t.Errorf("Failed to remove backend: %v", err)
	}
	if len(lb.ServerPool) != 0 {
		t.Errorf("Expected 0 backends, got %d", len(lb.ServerPool))
	}
}

func TestRemoveBackendNotFound(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	url1, _ := url.Parse("http://localhost:8001")
	backend := &server.Backend{
		Name: "NonExistent",
		Url:  url1,
	}

	err := lb.RemoveBackend(backend)
	if err == nil {
		t.Error("Expected error for non-existent backend, got nil")
	}
}

func TestHealthStatus(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	url1, _ := url.Parse("http://google.com")
	url2, _ := url.Parse("http://localhost:8002")

	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
	}
	backend1.Backend.Alive.Store(true)

	backend2 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server2",
			Url:  url2,
		},
	}
	backend2.Backend.Alive.Store(false)

	lb.AddBackend(backend1)
	lb.AddBackend(backend2)

	status := lb.HealthStatus()

	if len(status) != 2 {
		t.Errorf("Expected 2 backends in status, got %d", len(status))
	}
	if !status[backend1.Backend] {
		t.Error("Expected Server1 to be alive")
	}
	if status[backend2.Backend] {
		t.Error("Expected Server2 to be dead")
	}
}

func TestTrafficStatus(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	url1, _ := url.Parse("http://localhost:8001")
	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
		Balance: server.Serverbalance{
			Max_request: 100,
		},
	}
	backend1.Balance.Current_traffic.Store(50)
	backend1.Balance.Overalltraffic.Store(1000)

	lb.AddBackend(backend1)

	traffic := lb.TrafficStatus()

	if len(traffic) != 1 {
		t.Errorf("Expected 1 backend in traffic status, got %d", len(traffic))
	}
	if traffic[backend1.Backend]["current_traffic"] != 50 {
		t.Errorf("Expected current_traffic 50, got %d", traffic[backend1.Backend]["current_traffic"])
	}
	if traffic[backend1.Backend]["overall_traffic"] != 1000 {
		t.Errorf("Expected overall_traffic 1000, got %d", traffic[backend1.Backend]["overall_traffic"])
	}
}
