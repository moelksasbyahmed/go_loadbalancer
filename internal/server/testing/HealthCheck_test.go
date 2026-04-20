package LB_testing

import (
	"net/url"
	"testing"

	"github.com/moelksasbyahmed/go_loadbalancer/internal/server"
	"github.com/spf13/viper"
)

func TestRunManualCheckAll(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	url1, _ := url.Parse("http://localhost:8001")
	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
	}
	backend1.Backend.Alive.Store(false)

	lb.AddBackend(backend1)

	err := lb.RunManualCheck("all")
	if err != nil {
		t.Errorf("Unexpected error in RunManualCheck: %v", err)
	}
}

func TestRunManualCheckSpecificTarget(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	url1, _ := url.Parse("http://localhost:8001")
	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
	}
	backend1.Backend.Alive.Store(false)

	url2, _ := url.Parse("http://localhost:8002")
	backend2 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server2",
			Url:  url2,
		},
	}
	backend2.Backend.Alive.Store(false)

	lb.AddBackend(backend1)
	lb.AddBackend(backend2)

	// Check specific server by name
	err := lb.RunManualCheck("Server1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestRunManualCheckTargetNotFound(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	url1, _ := url.Parse("http://localhost:8001")
	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
	}
	lb.AddBackend(backend1)

	err := lb.RunManualCheck("NonExistent")
	if err == nil {
		t.Error("Expected error for target not found, got nil")
	}
}

func TestGetBackendStatus(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	url1, _ := url.Parse("http://localhost:8001")
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

	status := lb.GetBackendStatus()

	if len(status) != 2 {
		for backend, alive := range status {
			t.Logf("Backend: %s, Alive: %t", backend, alive)

		}
		t.Errorf("Expected 2 backends in status, got %d", len(status))

	}
	if !status["Server1"] {
		t.Error("Expected Server1 to be alive")
	}
	/*	if status["Server2"] { // wait till we have a backened that is working  to test this properly
		t.Error("Expected Server2 to be dead")
	}*/
}

func TestGetBackendStatusEmpty(t *testing.T) {
	lb := server.NewloadBalancer(&server.LoadBalancerConfig{}, viper.New())

	status := lb.GetBackendStatus()

	if len(status) != 0 {
		t.Errorf("Expected 0 backends in status, got %d", len(status))
	}
}
