BINARY_NAME=Go_LoadBalancer.exe


ifeq (run,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif


build:
	go build -o $(BINARY_NAME) ./cmd/CLI

run: clean  build
	./$(BINARY_NAME) $(RUN_ARGS)
	 

clean:
	go clean
	del $(BINARY_NAME)