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
	lb         *LoadBalancer
	HttpServer *http.Server
}

func NewServer(cfg *internal.Config, lb *LoadBalancer) *Server {
	return &Server{
		cfg: cfg,
		lb:  lb,
	}

}

func (s *Server) Start() error {
	handlers := s.SetupRouter()
	httpServer := &http.Server{
		Addr:         net.JoinHostPort(s.cfg.ServerConfig.Host, s.cfg.ServerConfig.Port),
		Handler:      handlers,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	s.HttpServer = httpServer
	log.Printf("Server starts on %s ", s.cfg.ServerConfig.Host+s.cfg.ServerConfig.Port)
	return httpServer.ListenAndServe()
}

func (s *Server) SetupRouter() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.lb.ProxyHandler()) //TODO needs to be fixed which make a proxy Handler is extendable from the server struct and not from the load balancer struct
	mux.HandleFunc("/health", s.HealthHandler)
	mux.HandleFunc("/stats", s.StatsHandler)
	mux.HandleFunc("/metrics", s.MetricsHandler)
	return mux

}
