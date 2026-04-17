package adminapi

import (
	"encoding/json"
	"net/http"

	"github.com/fatih/color"
)

func (api *AdminAPi) CheckHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	type payload struct {
		Name string `json:"name"`
		All  bool   `json:"all"`
	}
	var p payload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if p.All {
		api.LBServer.LB.RunManualCheck("all")
	} else {
		api.LBServer.LB.RunManualCheck(p.Name)
	}
	states := api.LBServer.LB.HealthStatus()
	response := make(map[string]bool)
	for backend, alive := range states {
		response[backend.Name] = alive
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	w.Write([]byte(color.GreenString("Health check completed successfully")))

}
