package server

import (
	"net/url"
	"sync/atomic"
)

type Backend struct {
	name  string
	url   *url.URL
	Alive atomic.Bool
}
