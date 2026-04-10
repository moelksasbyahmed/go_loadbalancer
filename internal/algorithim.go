package internal

import (
	server "github.com/moelkasabyahmed/go_loadbalancer/internal/server"
)

type Algorithim interface {
	nextserver() server.Backend
}
