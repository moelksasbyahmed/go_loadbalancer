package LB_testing

import (
	"net/url"
	"sync"
	"testing"

	"github.com/moelksasbyahmed/go_loadbalancer/internal/server"
)

func TestSetAlive(t *testing.T) {
	url1, _ := url.Parse("http://localhost:8001")
	backend := &server.Backend{
		Name: "TestServer",
		Url:  url1,
	}

	backend.SetAlive(true)
	if !backend.IsAlive() {
		t.Error("Expected backend to be alive")
	}

	backend.SetAlive(false)
	if backend.IsAlive() {
		t.Error("Expected backend to be dead")
	}
}

func TestIsAlive(t *testing.T) {
	url1, _ := url.Parse("http://localhost:8001")
	backend := &server.Backend{
		Name: "TestServer",
		Url:  url1,
	}

	backend.Alive.Store(true)
	if !backend.IsAlive() {
		t.Error("Expected backend to be alive")
	}

	backend.Alive.Store(false)
	if backend.IsAlive() {
		t.Error("Expected backend to be dead")
	}
}

func TestBackendConcurrentReadWrite(t *testing.T) {
	url1, _ := url.Parse("http://localhost:8001")
	backend := &server.Backend{
		Name: "TestServer",
		Url:  url1,
	}
	backend.Alive.Store(true)

	done := make(chan bool)
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(state bool) {
			defer wg.Done()
			backend.SetAlive(state)
		}(i%2 == 0)
	}

	// Concurrent reads
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = backend.IsAlive()
		}()
	}

	go func() {
		wg.Wait()
		done <- true
	}()

	<-done
}

func TestBackendHealthCheckFailure(t *testing.T) {
	url1, _ := url.Parse("http://localhost:9999") // Non-existent server
	backend := &server.Backend{
		Name: "TestServer",
		Url:  url1,
	}
	backend.Alive.Store(true)

	backend.HealthCheck()

	// Should be marked as dead after failed health check
	if backend.IsAlive() {
		t.Error("Expected backend to be marked as dead after failed health check")
	}
}
