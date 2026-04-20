package server

import (
	"fmt"
	"strings"
	"sync"
)

var AlgorithimsRegistry = make(map[string]func() BalancerAlgorithm)

type BalancerAlgorithm interface {
	NextPeer(backends []*LoadBalancerUnit) (*Backend, error)
}

func RegisterAlgorithim(name string, Algorithim func() BalancerAlgorithm) {
	AlgorithimsRegistry[name] = Algorithim

}
func GetAlgorithim(name string) (BalancerAlgorithm, error) {
	name = strings.ToLower(name)
	fmt.Println(name)
	Alg, exist := AlgorithimsRegistry[name]
	if !exist {
		return nil, fmt.Errorf(" %s is UNKNOWN algorithim ", name)

	}
	return Alg(), nil
}
func GetAVailableAlgorithims() []string {
	algorithims := make([]string, 0, len(AlgorithimsRegistry))
	for name := range AlgorithimsRegistry {
		algorithims = append(algorithims, name)
	}
	return algorithims

}

type RoundRobin struct {
	mux     sync.RWMutex
	current int
}

func (rr *RoundRobin) NextPeer(backends []*LoadBalancerUnit) (*Backend, error) {
	rr.mux.Lock()
	defer rr.mux.Unlock()

	if len(backends) == 0 {
		return nil, fmt.Errorf("no backends available")
	}
	for i := 0; i < len(backends); i++ {
		rr.current = (rr.current + 1) % len(backends)
		candidate := backends[rr.current]
		if !candidate.Backend.Alive.Load() {
			continue
		}
		if candidate.Balance.Current_traffic.Load() >= int64(candidate.Balance.Max_request) {
			continue
		}
		return candidate.Backend, nil
	}
	return nil, fmt.Errorf("no backends available")

}

func init() {
	RegisterAlgorithim("round_robin", func() BalancerAlgorithm {
		return &RoundRobin{}
	})
}
