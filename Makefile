# Name of the binary
BINARY_NAME=matrixterm

# Build output directory
BUILD_DIR=build

# Go source files
SRC=main.go helpers.go

# Default target
all: build

# Build the binary into the build directory
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC)

# Run the binary
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)
