package server

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

func (lb *LoadBalancer) ProxyHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		Backend, err := lb.config.Algorithim.NextPeer(lb.ServerPool)
		for _, servers := range lb.ServerPool {
			log.Printf("Backend %s is alive: %t, current traffic: %d, max request: %d\n", servers.Backend.Url, servers.Backend.Alive.Load(), servers.Balance.current_traffic.Load(), servers.Balance.Max_request)
		}
		if Backend == nil {
			log.Printf("Backend is nil - no available servers")
			http.Error(w, "No backends available", http.StatusServiceUnavailable)
			return
		}
		if err != nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}
		Director := func(req *http.Request) {

			req.URL.Scheme = Backend.Url.Scheme
			req.URL.Host = Backend.Url.Host
			req.Host = Backend.Url.Host
			req.Header.Set("X-Proxy-Timestamp", time.Now().UTC().Format(time.RFC3339))
			clientIP := req.RemoteAddr
			if forwardedFor := req.Header.Get("X-Forwarded-For"); forwardedFor != "" {
				clientIP = forwardedFor + ", " + clientIP
			}
			req.Header.Set("X-Forwarded-For", clientIP)
			req.Header.Set("X-Forwarded-Host", req.Host)
			req.Header.Set("X-Forwarded-Proto", req.URL.Scheme)
			//log.Printf("proxying %s  to %s  (Client_ip %s)\n ", req.URL.Path, Backend.Url, clientIP)

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
		for _, servers := range lb.ServerPool {
			if servers.Backend.Name == Backend.Name && servers.Backend.Url.String() == Backend.Url.String() {
				servers.Balance.current_traffic.Add(1)
				servers.Balance.overalltraffic.Add(1)
				defer servers.Balance.current_traffic.Add(-1)
				break
			}
		}

		proxy.ServeHTTP(w, r)

	}
}
