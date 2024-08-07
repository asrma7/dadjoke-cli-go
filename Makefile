
# Detect current OS and architecture
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
GOBIN := $(shell go env GOBIN)
GOPATH := $(shell go env GOPATH)

# Define output file name
OUT_FILE := dadjoke
ifeq ($(GOOS),windows)
	OUT_FILE := dadjoke.exe
endif

# Build executable for current OS and architecture
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUT_FILE)

# Install the executable for current OS and architecture
install: build
	@if [ "$(GOOS)" = "windows" ]; then \
		if [ -n "$(GOBIN)" ]; then \
			mkdir -p "$(GOBIN)"; \
			cp $(OUT_FILE) "$(GOBIN)\\$(OUT_FILE)"; \
		elif [ -n "$(GOPATH)" ]; then \
			mkdir -p "$(GOPATH)\\bin"; \
			cp $(OUT_FILE) "$(GOPATH)\\bin\\$(OUT_FILE)"; \
		else \
			echo "GOBIN or GOPATH is not set. Please set GOBIN or GOPATH to the installation directory."; \
			exit 1; \
		fi \
	else \
		if [ -n "$(GOBIN)" ]; then \
			mkdir -p "$(GOBIN)"; \
			install -m 755 $(OUT_FILE) "$(GOBIN)/$(OUT_FILE)"; \
		elif [ -n "$(GOPATH)" ]; then \
			mkdir -p "$(GOPATH)/bin"; \
			install -m 755 $(OUT_FILE) "$(GOPATH)/bin/$(OUT_FILE)"; \
		else \
			echo "GOBIN or GOPATH is not set. Please set GOBIN or GOPATH to the installation directory."; \
			exit 1; \
		fi \
	fi

# Clean the build
clean:
	rm -f $(OUT_FILE)

.PHONY: build install clean
