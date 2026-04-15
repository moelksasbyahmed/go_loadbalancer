package adminapi

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/moelksasbyahmed/go_loadbalancer/internal"
	"github.com/moelksasbyahmed/go_loadbalancer/internal/server"
)

type AdminAPi struct {
	lb     *server.LoadBalancer
	server *http.Server
	config *internal.Config
}

func NewAdminAPI(lb *server.LoadBalancer, config *internal.Config) *AdminAPi {
	return &AdminAPi{
		lb:     lb,
		config: config,
	}
}

func (api *AdminAPi) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/check", api.CheckHandler)
	mux.HandleFunc("/add", api.AddHandler)
	mux.HandleFunc("/remove", api.handleRemoveServer)
	mux.HandleFunc("/status", api.StatusHandler)
	mux.HandleFunc("/list", api.ListHandler)
	api.server = &http.Server{
		Addr:         net.JoinHostPort(api.config.AdminConfig.Host, api.config.AdminConfig.Port),
		Handler:      mux,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Printf("Starting Admin Server on http://%s:%s ", api.config.AdminConfig.Host, api.config.AdminConfig.Port)
	return api.server.ListenAndServe()

}
