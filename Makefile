BINARY_NAME=valueanalyzer

IMAGE_NAME=valueanalyzer

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	GOOS = linux
	GOARCH = amd64
endif
ifeq ($(UNAME_S),Darwin)
	GOOS = darwin
	GOARCH = amd64
endif

# Build the Go application
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY_NAME) .

# Run the Go application
run:
	./$(BINARY_NAME)

# Clean the build files
clean:
	rm -f $(BINARY_NAME)

# Docker build
docker-build:
	docker build -t $(IMAGE_NAME) .

# Docker run
docker-run:
	docker run -e API_KEY=$(API_KEY) -e TELEGRAM_BOT_TOKEN=$(TELEGRAM_BOT_TOKEN) -e TELEGRAM_CHAT_ID=$(TELEGRAM_CHAT_ID) $(IMAGE_NAME)

# Docker clean
docker-clean:
	docker rmi $(IMAGE_NAME)

.PHONY: build run clean docker-build docker-run docker-clean
