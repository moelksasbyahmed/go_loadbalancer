package server

import (
	"log"
	"net"
	"net/http"
	"time"

	internal "github.com/moelksasbyahmed/go_loadbalancer/internal"
)

type Server struct {
	cfg        *internal.Config
	LB         *LoadBalancer
	HttpServer *http.Server
}

func NewServer(cfg *internal.Config, lb *LoadBalancer) *Server {
	return &Server{
		cfg: cfg,
		LB:  lb,
	}

}

func (s *Server) Start() error {
	handlers := s.SetupRouter()
	httpServer := &http.Server{
		Addr:         net.JoinHostPort(s.cfg.LoadBalancerConfig.Host, s.cfg.LoadBalancerConfig.Port),
		Handler:      handlers,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	s.HttpServer = httpServer
	log.Printf("Server starts on http://%s:%s ", s.cfg.LoadBalancerConfig.Host, s.cfg.LoadBalancerConfig.Port)
	return httpServer.ListenAndServe()
}

func (s *Server) SetupRouter() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.LB.ProxyHandler()) //TODO needs to be fixed which make a proxy Handler is extendable from the server struct and not from the load balancer struct
	mux.HandleFunc("/health", s.HealthHandler)
	mux.HandleFunc("/stats", s.StatsHandler)
	mux.HandleFunc("/metrics", s.MetricsHandler) // OLD DESIGN BUT we will keep it for now until we decide if we will try to implement  these handlers or leave the Excution Flow in the Admin api as it is for now Data 4/19/2026
	return mux

}

func (s *Server) CloseServer() error {
	return s.HttpServer.Close()
}
