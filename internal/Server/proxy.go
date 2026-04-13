package server

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

func StartLoadBalancer(endpoint string, port string, Algorithim BalancerAlgorithm) error {
	LoadBalancer := NewloadBalancer(Algorithim)
	proxyHandler := ProxyHandler(LoadBalancer)

	server := &http.Server{
		Addr:         net.JoinHostPort(endpoint, port),
		Handler:      proxyHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Printf("Starting load balancer on %s:%s\n", endpoint, port)
	err := server.ListenAndServe()
	if err != nil {
		return err

	}

	return nil
}
func ProxyHandler(lb *LoadBalancer) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		Backend, err := lb.Algorithim.NextPeer(lb.ServerPool)
		if err != nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		}
		Director := func(req *http.Request) {

			req.URL.Scheme = Backend.url.Scheme
			req.URL.Host = Backend.url.Host
			req.Host = Backend.url.Host
			req.Header.Set("X-Proxy-Timestamp", time.Now().UTC().Format(time.RFC3339))
			clientIP := req.RemoteAddr
			if forwardedFor := req.Header.Get("X-Forwarded-For"); forwardedFor != "" {
				clientIP = forwardedFor + ", " + clientIP
			}
			req.Header.Set("X-Forwarded-For", clientIP)
			req.Header.Set("X-Forwarded-Host", req.Host)
			req.Header.Set("X-Forwarded-Proto", req.URL.Scheme)
			log.Printf("proxying %s  to %s  (Client_ip %s)\n ", req.URL.Path, Backend.url, clientIP)

		}
		Transport := &http.Transport{

			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 40,
			MaxConnsPerHost:     25,

			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 60 * time.Second,
			}).DialContext,
		}

		proxy := &httputil.ReverseProxy{
			Director:  Director,
			Transport: Transport,
			ErrorHandler: func(rw http.ResponseWriter, req *http.Request, err error) {
				log.Printf("Proxy error: %v", err)
				http.Error(rw, "Bad Gateway", http.StatusBadGateway)
			},
		}

		proxy.ServeHTTP(w, r)

	}
}
