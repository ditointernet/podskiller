# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=podskiller
BINARY_DEV=$(BINARY_NAME)_dev

all: test build
build:
	$(GOBUILD) -o $(BINARY_DEV) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_DEV)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_DEV) -v ./...
	./$(BINARY_DEV)
build-container:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o $(BINARY_NAME) -v .
	docker build -t ditointernet/podskiller:latest .
