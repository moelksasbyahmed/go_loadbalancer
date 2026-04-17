package server

import (
	"net/http"
	"time"
)

func (back *Backend) SetAlive(alive bool) {
	back.mux.Lock()
	back.Alive.Store(alive)
	defer back.mux.Unlock()

}

func (back *Backend) IsAlive() bool {
	back.mux.RLock()
	defer back.mux.RUnlock()
	return back.Alive.Load()
}

func (back *Backend) HealthCheck() {
	ping_url := back.Url.String() + "/"
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(ping_url)
	if err != nil || resp.StatusCode != http.StatusOK {
		back.SetAlive(false)

		return
	}

	back.SetAlive(true)
}
