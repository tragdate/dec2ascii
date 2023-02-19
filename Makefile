# Define the name of the executable
EXECUTABLE=dec2ascii

# Define the source files
SRC=main.go

# Define the output directory
OUTPUT_DIR=bin

# Define the installation directory
INSTALL_DIR=/usr/local/bin

# Define the compiler
GO=go

# Define the build flags
BUILD_FLAGS=-v

# Define the install flags
INSTALL_FLAGS=

# Define the clean flags
CLEAN_FLAGS=

# Build the executable
build:
	@echo "Building $(EXECUTABLE)..."
	@mkdir -p $(OUTPUT_DIR)
	@$(GO) build $(BUILD_FLAGS) -o $(OUTPUT_DIR)/$(EXECUTABLE) $(SRC)

# Install the executable
install: build
	@echo "Installing $(EXECUTABLE)..."
	@mkdir -p $(INSTALL_DIR)
	@cp $(OUTPUT_DIR)/$(EXECUTABLE) $(INSTALL_DIR)/$(EXECUTABLE) $(INSTALL_FLAGS)

# Clean the build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(OUTPUT_DIR) $(CLEAN_FLAGS)
