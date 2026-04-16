package adminapi

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/moelksasbyahmed/go_loadbalancer/internal"
	LB "github.com/moelksasbyahmed/go_loadbalancer/internal/server"
)

type AdminAPi struct {
	LBServer *LB.Server
	server   *http.Server
	config   *internal.Config
}

func NewAdminAPI(Lbserver *LB.Server, config *internal.Config) *AdminAPi {
	return &AdminAPi{
		LBServer: Lbserver,
		config:   config,
	}
}

func (api *AdminAPi) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.LoggerHandler)
	mux.HandleFunc("/check", api.CheckHandler)
	mux.HandleFunc("/add", api.AddHandler)
	mux.HandleFunc("/remove", api.handleRemoveServer)
	mux.HandleFunc("/status", api.StatusHandler)
	mux.HandleFunc("/list", api.ListHandler)
	mux.HandleFunc("/abort", api.AbortHandler)
	api.server = &http.Server{
		Addr:         net.JoinHostPort(api.config.Adminconfig.Host, api.config.Adminconfig.Port),
		Handler:      mux,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Printf("Starting Admin Server on http://%s:%s ", api.config.Adminconfig.Host, api.config.Adminconfig.Port)
	return api.server.ListenAndServe()

}
