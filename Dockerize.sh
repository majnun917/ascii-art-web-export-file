#!/bin/sh

# Variables
IMAGE_NAME="doc-image"
IMAGE_TAG="1.0"
CONTAINER_NAME="doc-container"
PORT=8080

# Build the Docker image
echo "Building Docker image..."
docker build -t $IMAGE_NAME:$IMAGE_TAG .

# Check if build was successful
if [ $? -ne 0 ]; then
    echo "Docker image build failed!"
    exit 1
fi

# Stop and remove any existing container with the same name
echo "Stopping and removing any existing container..."
docker stop $CONTAINER_NAME
docker rm $CONTAINER_NAME

# Run the Docker container
echo "Running Docker container..."
docker run -d -p $PORT:$PORT --name $CONTAINER_NAME $IMAGE_NAME:$IMAGE_TAG

# Check if the container started successfully
if [ $? -ne 0 ]; then
    echo "Failed to start Docker container!"
    exit 1
fi

echo "Docker container is running at localhost:$PORT"