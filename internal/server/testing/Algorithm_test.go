package LB_testing

import (
	"net/url"
	"testing"

	"github.com/moelksasbyahmed/go_loadbalancer/internal/server"
)

func TestRegisterAlgorithm(t *testing.T) {
	customAlg := func() server.BalancerAlgorithm {
		return &server.RoundRobin{}
	}

	server.RegisterAlgorithim("test_custom", customAlg)

	alg, err := server.GetAlgorithim("test_custom")
	if err != nil {
		t.Errorf("Failed to get registered algorithm: %v", err)
	}
	if alg == nil {
		t.Error("Expected algorithm, got nil")
	}
}

func TestGetAlgorithmNotFound(t *testing.T) {
	alg, err := server.GetAlgorithim("nonexistent_algorithm")
	if err == nil {
		t.Error("Expected error for non-existent algorithm, got nil")
	}
	if alg != nil {
		t.Error("Expected nil algorithm, got value")
	}
}

func TestGetAvailableAlgorithms(t *testing.T) {
	algs := server.GetAVailableAlgorithims()
	if len(algs) == 0 {
		t.Error("Expected at least one algorithm registered")
	}

	found := false
	for _, alg := range algs {
		if alg == "round_robin" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected 'round_robin' algorithm in available algorithms")
	}
}

func TestRoundRobinNextPeer(t *testing.T) {
	rr := &server.RoundRobin{}

	url1, _ := url.Parse("http://localhost:8001")
	url2, _ := url.Parse("http://localhost:8002")

	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
		Balance: server.Serverbalance{
			Max_request: 100,
		},
	}
	backend1.Backend.Alive.Store(true)

	backend2 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server2",
			Url:  url2,
		},
		Balance: server.Serverbalance{
			Max_request: 100,
		},
	}
	backend2.Backend.Alive.Store(true)

	backends := []*server.LoadBalancerUnit{backend1, backend2}

	peer1, err := rr.NextPeer(backends)
	if err != nil {
		t.Errorf("Failed to get next peer: %v", err)
	}
	if peer1 == nil {
		t.Error("Expected peer, got nil")
	}

	peer2, err := rr.NextPeer(backends)
	if err != nil {
		t.Errorf("Failed to get next peer: %v", err)
	}
	if peer2 == nil {
		t.Error("Expected peer, got nil")
	}
}

func TestRoundRobinNoBackendsAvailable(t *testing.T) {
	rr := &server.RoundRobin{}
	backends := []*server.LoadBalancerUnit{}

	_, err := rr.NextPeer(backends)
	if err == nil {
		t.Error("Expected error for no backends available, got nil")
	}
}

func TestRoundRobinAllBackendsDead(t *testing.T) {
	rr := &server.RoundRobin{}

	url1, _ := url.Parse("http://localhost:8001")
	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
	}
	backend1.Backend.Alive.Store(false)

	backends := []*server.LoadBalancerUnit{backend1}

	_, err := rr.NextPeer(backends)
	if err == nil {
		t.Error("Expected error when all backends are dead, got nil")
	}
}

func TestRoundRobinTrafficLimit(t *testing.T) {
	rr := &server.RoundRobin{}

	url1, _ := url.Parse("http://localhost:8001")
	url2, _ := url.Parse("http://localhost:8002")

	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server1",
			Url:  url1,
		},
		Balance: server.Serverbalance{
			Max_request: 10,
		},
	}
	backend1.Backend.Alive.Store(true)
	backend1.Balance.Current_traffic.Store(10) // At capacity

	backend2 := &server.LoadBalancerUnit{
		Backend: &server.Backend{
			Name: "Server2",
			Url:  url2,
		},
		Balance: server.Serverbalance{
			Max_request: 100,
		},
	}
	backend2.Backend.Alive.Store(true)

	backends := []*server.LoadBalancerUnit{backend1, backend2}

	peer, err := rr.NextPeer(backends)
	if err != nil {
		t.Errorf("Failed to get next peer: %v", err)
	}
	if peer != backend2.Backend {
		t.Error("Expected to get Server2 since Server1 is at capacity")
	}
}

func TestRoundRobinDistribution(t *testing.T) {
	rr := &server.RoundRobin{}

	url1, _ := url.Parse("http://localhost:8001")
	url2, _ := url.Parse("http://localhost:8002")
	url3, _ := url.Parse("http://localhost:8003")

	backend1 := &server.LoadBalancerUnit{
		Backend: &server.Backend{Name: "Server1", Url: url1},
		Balance: server.Serverbalance{Max_request: 100},
	}
	backend1.Backend.Alive.Store(true)

	backend2 := &server.LoadBalancerUnit{
		Backend: &server.Backend{Name: "Server2", Url: url2},
		Balance: server.Serverbalance{Max_request: 100},
	}
	backend2.Backend.Alive.Store(true)

	backend3 := &server.LoadBalancerUnit{
		Backend: &server.Backend{Name: "Server3", Url: url3},
		Balance: server.Serverbalance{Max_request: 100},
	}
	backend3.Backend.Alive.Store(true)

	backends := []*server.LoadBalancerUnit{backend1, backend2, backend3}

	counts := make(map[string]int)
	for i := 0; i < 9; i++ {
		peer, _ := rr.NextPeer(backends)
		counts[peer.Name]++
	}

	if counts["Server1"] != 3 || counts["Server2"] != 3 || counts["Server3"] != 3 {
		t.Errorf("Expected balanced distribution, got %v", counts)
	}
}
