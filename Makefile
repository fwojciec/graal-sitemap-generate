GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run
BINARY_NAME=graal_sitemap_generate
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: build build-linux clean 

build:
	$(GOBUILD) -o $(BINARY_NAME) main.go sitemap.go slugs.go

build-linux:
	env GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) main.go sitemap.go slugs.go

clean:
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)