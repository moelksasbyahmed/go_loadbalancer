package adminapi

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/fatih/color"
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

	corsWrapper := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		mux.ServeHTTP(w, r)
	})

	mux.HandleFunc("/logger", api.LoggerHandler)
	mux.HandleFunc("/check", api.CheckHandler)
	mux.HandleFunc("/add", api.AddHandler)
	mux.HandleFunc("/remove", api.handleRemoveServer)
	mux.HandleFunc("/status", api.StatusHandler)
	mux.HandleFunc("/list", api.ListHandler)
	mux.HandleFunc("/abort", api.AbortHandler)
	api.server = &http.Server{
		Addr:         net.JoinHostPort(api.config.Adminconfig.Host, api.config.Adminconfig.Port),
		Handler:      corsWrapper,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Printf("Starting Admin Server on http://%s:%s ", api.config.Adminconfig.Host, api.config.Adminconfig.Port)
	return api.server.ListenAndServe()

}

func (api *AdminAPi) Shutdown(ctx context.Context) error {
	log.Println(color.YellowString("Shutting down the Load Balancer..."))
	if err := api.LBServer.HttpServer.Shutdown(ctx); err != nil {
		log.Printf("LB Server shutdown error: %v\n", err)
	}

	log.Println(color.YellowString("Shutting down the Admin API..."))
	if err := api.server.Shutdown(ctx); err != nil {
		log.Printf("Admin API shutdown error: %v\n", err)
	}

	return nil
}
