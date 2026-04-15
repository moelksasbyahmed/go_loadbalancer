package adminapi

import (
	"net/http"

	_ "github.com/moelksasbyahmed/go_loadbalancer/internal/server"
)

func (api *AdminAPi) AbortHandler(w http.ResponseWriter, r *http.Request) {

	err := api.LBServer.CloseServer()
	if err != nil {
		http.Error(w, "Failed to abort the LoadBalancer server ", http.StatusInternalServerError)
		return
	}
	/*
		er := api.server.Close()
		if er != nil {
			http.Error(w, "Failed to abort the Admin server ", http.StatusInternalServerError)
			return
		}
	*/
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server aborted successfully"))

}
