DOCKER_IMAGE_NAME := go-web-page-fetcher
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
	@docker run -v $(OUTPUT_DIR):/output $(DOCKER_IMAGE_NAME) --output /output $(ARGS)

clean:
	@echo "Cleaning up..."
	go clean
	# Stop and remove all containers associated with the image
	-docker ps -a -q --filter ancestor=$(DOCKER_IMAGE_NAME) | xargs -r docker rm -f
	# Remove the image
	-docker rmi $(DOCKER_IMAGE_NAME)

.PHONY: build docker-build docker-run clean
