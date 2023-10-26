DOCKER_IMAGE_NAME := go-web-page-fetcher
CONTAINER_NAME := go-web-page-fether-container
OUTPUT_DIR := .

ARGS :=

# Make targets
build:
	@echo "Building Go application..."
	go mod download
	go build .

docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	@echo "Running fetch tool inside Docker container..."
	@CONTAINER_REMOVED=$$(docker rm -f $(CONTAINER_NAME) 2>/dev/null); \
	if [ -n "$$CONTAINER_REMOVED" ]; then \
		echo "Removed existing container: $(CONTAINER_NAME), creating new one."; \
	else \
		echo "No container named $(CONTAINER_NAME) to remove, creating new one."; \
	fi
	@docker run --name $(CONTAINER_NAME) -v $(OUTPUT_DIR):/output $(DOCKER_IMAGE_NAME) --output /output $(ARGS)

clean:
	@echo "Cleaning up..."
	-go clean
	# Stop and remove all containers associated with the image
	-docker ps -a -q --filter ancestor=$(DOCKER_IMAGE_NAME) | xargs -r docker rm -f
	# Remove the image
	-docker rmi $(DOCKER_IMAGE_NAME)

.PHONY: build docker-build docker-run clean
