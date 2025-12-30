# Makefile for FastPM (Go core + TypeScript CLI)

# Paths
GO_SRC := src/main.go
GO_BUILD := build/fastpm.exe
TS_SRC := tsconfig.json
TS_BUILD := dist/cli.js

# Default target
.PHONY: all
all: build

# -------------------------------
# Build Go binary
# -------------------------------
.PHONY: build-go
build-go:
	@echo "🔹 Building Go core..."
	@mkdir -p build
	go build -o $(GO_BUILD) $(GO_SRC)
	@echo "✅ Go core built: $(GO_BUILD)"

# -------------------------------
# Compile TypeScript CLI
# -------------------------------
.PHONY: build-ts
build-ts:
	@echo "🔹 Compiling TypeScript CLI..."
	tsc
	@echo "✅ TypeScript CLI compiled: $(TS_BUILD)"

# -------------------------------
# Full build (Go + TS)
# -------------------------------
.PHONY: build
build: build-go build-ts
	@echo "🚀 FastPM build complete!"

# -------------------------------
# Clean build artifacts
# -------------------------------
.PHONY: clean
clean:
	@echo "🧹 Cleaning build and dist folders..."
	rm -rf build dist
	@echo "✅ Cleaned!"

# -------------------------------
# Link CLI globally (npm link)
# -------------------------------
.PHONY: link
link:
	@echo "🔗 Linking FastPM CLI globally..."
	npm link
	@echo "✅ FastPM CLI linked globally!"

# -------------------------------
# Test FastPM CLI locally
# -------------------------------
.PHONY: test
test:
	@echo "🧪 Running FastPM test..."
	node $(TS_BUILD) --help

# -------------------------------
# Install a package using FastPM
# Usage: make install PACKAGE=eslint
# -------------------------------
.PHONY: install
install:
ifndef PACKAGE
	$(error PACKAGE is not set. Example: make install PACKAGE=eslint)
endif
	@echo "📦 Installing package: $(PACKAGE)"
	node $(TS_BUILD) install $(PACKAGE)

# -------------------------------
# Shortcut: rebuild + link + test
# -------------------------------
.PHONY: rebuild
rebuild: clean build link test
	@echo "🔧 Rebuild + link + test complete!"
