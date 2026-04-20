package server

import "net/http"

func (s *Server) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
"status" : "healthy"
}"`))

}
