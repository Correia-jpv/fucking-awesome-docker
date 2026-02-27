BINARY := awesome-docker
.PHONY: build test lint check health report clean

build:
	go build -o $(BINARY) ./cmd/awesome-docker

test:
	go test ./internal/... -v

lint: build
	./$(BINARY) lint

check: build
	./$(BINARY) check

health: build
	./$(BINARY) health

report: build
	./$(BINARY) report

clean:
	rm -f $(BINARY)
