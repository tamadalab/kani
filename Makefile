GO := go
VERSION := 1.0.0
NAME := bin/kani
DIST := $(NAME)-$(VERSION)

all: $(NAME)

$(NAME): git-kani.go cmd/enableCmd.go cmd/initCmd.go cmd/rootCmd.go cmd/storeCmd.go cmd/runAnalysisEngines.go
	$(GO) build -o $(NAME) git-kani.go

define _createDist
    mkdir -p dist/$(1)_$(2)/$(DIST)
    GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME) cmd/$(NAME)/*.go
    cp -r README.md LICENSE completions dist/$(1)_$(2)/$(DIST)
    cp -r docs/public dist/$(1)_$(2)/$(DIST)/docs
    tar cfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
endef

dist: build docs
	@$(call _createDist,darwin,amd64)
	@$(call _createDist,darwin,arm64)
	@$(call _createDist,windows,amd64)
	@$(call _createDist,windows,386)
	@$(call _createDist,linux,amd64)
	@$(call _createDist,linux,386)

clean:
	$(GO) clean
	rm -rf $(NAME) dist
