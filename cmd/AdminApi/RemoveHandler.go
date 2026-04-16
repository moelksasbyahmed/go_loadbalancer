package adminapi

import (
	"encoding/json"
	"fmt"

	"net/http"
	"net/url"

	"github.com/moelksasbyahmed/go_loadbalancer/internal/server"
)

func (api *AdminAPi) handleRemoveServer(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
	var back backend
	err := json.NewDecoder(r.Body).Decode(&back)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
	}
	RMV_URL, _ := url.Parse(back.Url)
	fmt.Println(back.Name)
	err = api.LBServer.LB.RemoveBackend(&server.Backend{
		Name: back.Name,
		Url:  RMV_URL,
	})
	if err != nil {
		http.Error(w, "Server Doesnt Exist", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server removed successfully"))
}
