package adminapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (api *AdminAPi) ListHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	servers := api.LBServer.LB.HealthStatus()
	type response struct {
		Name  string  `json:"name"`
		Alive bool    `json:"alive"`
		Url   url.URL `json:"url"`
	}
	var res []response
	for backend, alive := range servers {
		res = append(res, response{
			Name:  backend.Name,
			Url:   *backend.Url,
			Alive: alive,
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
