package server

type BalancerAlgorithm interface {
	NextPeer(backends []*LoadBalancerUnit) (*Backend, error)
}
type RoundRobin struct {
	current int
}

func (rr *RoundRobin) NextPeer(backends []*Backend) *Backend {

	return nil
}
