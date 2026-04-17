package adminapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/moelksasbyahmed/go_loadbalancer/internal"
	server "github.com/moelksasbyahmed/go_loadbalancer/internal/server"
	"github.com/spf13/viper"
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
	var currentServers []internal.Serversconfig
	if err := viper.UnmarshalKey("servers", &currentServers); err != nil {
		log.Printf("Failed to unmarshal servers config: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	currentServers = append(currentServers, internal.Serversconfig{
		Name:       b.Name,
		URl:        b.Url,
		Alive:      true,
		MaxRequest: b.MaxRequest,
	})
	viper.Set("servers", currentServers)

	
	if err := viper.WriteConfig(); err != nil {
		log.Printf("Failed to write config file: %v", err)
		http.Error(w, "Failed to save configuration", http.StatusInternalServerError)
		return
	}
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
