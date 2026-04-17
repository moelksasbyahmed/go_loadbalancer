package adminapi

import (
	"encoding/json"
	"net/http"
)

func (api *AdminAPi) StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	states := api.LBServer.LB.TrafficStatus()
	type response struct {
		Name           string `json:"name"`
		Alive          bool   `json:"alive"`
		Url            string `json:"url"`
		CurrentTraffic int    `json:"current_traffic"`
		OverallTraffic int    `json:"overall_traffic"`
	}
	var res []response
	for backend, traffic := range states {
		res = append(res, response{
			Name:           backend.Name,
			Alive:          backend.Alive.Load(),
			Url:            backend.Url.String(),
			CurrentTraffic: int(traffic["current_traffic"]),
			OverallTraffic: int(traffic["overall_traffic"]),
		})
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
