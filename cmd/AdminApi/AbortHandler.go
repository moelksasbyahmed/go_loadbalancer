package adminapi

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
	_ "github.com/moelksasbyahmed/go_loadbalancer/internal/server"
)

func (api *AdminAPi) AbortHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Println(color.RedString("Abort command received via API"))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Shutdown initiated\n"))

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		api.Shutdown(ctx)
	}()

}
