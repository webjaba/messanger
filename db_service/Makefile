BINARY_NAME = bin/db_service
MAIN_FILE = cmd/main.go
BUILD_FLAGS = -v
CFG = ./config/local.yaml

build:
	@echo "Building..."
    CONFIG_PATH=$(CFG) go build $(BUILD_FLAGS) -o $(BINARY_NAME) $(MAIN_FILE)

run:
	@echo "Running..."
	CONFIG_PATH=$(CFG) go run $(MAIN_FILE)