package adminapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	server "github.com/moelksasbyahmed/go_loadbalancer/internal/server"
)

func (api *AdminAPi) AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var b backend
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	back, err := Payload2Backend(&b)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse backend: %v", err), http.StatusBadRequest)
		return
	}
	er := api.LBServer.LB.AddBackend(back)
	if er != nil {
		http.Error(w, fmt.Sprintf("Failed to add backend: %v", err), http.StatusBadRequest)
		return
	}
	api.LBServer.LB.ViperSync()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("the backend added successfully "))
}

func Payload2Backend(b *backend) (*server.LoadBalancerUnit, error) {
	parsedURL, err := url.Parse(b.Url)
	if err != nil {
		return nil, fmt.Errorf("Invalid URL: %v", err)
	}
	bkserver := &server.Backend{
		Name: b.Name,
		Url:  parsedURL,
	}

	return &server.LoadBalancerUnit{
		Backend: bkserver,
		Balance: server.Serverbalance{
			Max_request: b.MaxRequest},
	}, nil
}
