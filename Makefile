BINARY_NAME=Go_LoadBalancer.exe


ifeq (run,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

.PHONY: build run clean run-containers stop clean-containers restart status help


build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) ./cmd/CLI
	@echo "Build complete"


run: clean build
	@echo "Starting Go_LoadBalancer..."
	./$(BINARY_NAME) $(RUN_ARGS)


clean:
	@echo "Cleaning build artifacts..."
	go clean
	-del $(BINARY_NAME) 2>nul || rm -f $(BINARY_NAME)
	@echo " Clean complete"

run-containers:
	@echo "Starting mock backend servers..."
	@echo "  If containers already exist, run 'make stop' first"
	@docker run --rm -d -p 8001:80 --name server1 kennethreitz/httpbin
	@docker run --rm -d -p 8002:80 --name server2 kennethreitz/httpbin
	@docker run --rm -d -p 8003:80 --name server3 kennethreitz/httpbin
	@echo " All containers started successfully!"
	@echo ""
	@docker ps --filter "name=server" --format "table {{.Names}}\t{{.Ports}}\t{{.Status}}"

stop:
	@echo "Stopping mock backend servers..."
	@docker stop server1 server2 server3 2>nul 2>/dev/null || echo "Containers already stopped"
	@echo "All containers stopped"


clean-containers: stop
	@echo "Removing containers..."
	@-docker rm server1 server2 server3 2>nul 2>/dev/null || true
	@echo " Containers removed"


restart: clean-containers run-containers


status:
	@echo "Container Status:"
	@docker ps -a --filter "name=server" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" || echo "No containers found"


test:
	@echo "Testing backend servers..."
	@for i in 1 2 3; do \
		echo -n "Server$$i on port 800$$i: "; \
		curl -s -o /dev/null -w "%{http_code}" http://localhost:800$$i/get && echo ""; \
	done


help:
	@echo "Available commands:"
	@echo "  make build           - Build the Go application"
	@echo "  make run [args]      - Build and run the application"
	@echo "  make clean           - Clean build artifacts"
	@echo "  make run-containers  - Start mock backend Docker containers"
	@echo "  make stop            - Stop mock backend containers"
	@echo "  make clean-containers- Stop and remove containers"
	@echo "  make restart         - Restart all containers"
	@echo "  make status          - Show container status"
	@echo "  make test            - Test backend connectivity"
	@echo "  make help            - Show this help message"


.DEFAULT_GOAL := help