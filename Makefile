BIN_DIR := $(or $(BIN_DIR), ./build)
BINARY_NAME := $(or $(BIN_NAME), api)

clean:
	rm -rf "$(BIN_DIR)/*"

build:
	mkdir -p ./$(BIN_DIR) && go build -o "$(BIN_DIR)/$(BINARY_NAME)" ./cmd/main.go

run:
	go run ./cmd/main.go
