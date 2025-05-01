#!/bin/bash

# Display script banner
echo "====================================================="
echo "      Starting Application Docker Containers"
echo "====================================================="

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
  echo "Error: Docker is not running. Please start Docker and try again."
  exit 1
fi

# Stop and remove any existing containers with the same names
echo "Cleaning up any existing containers..."
docker-compose down 2>/dev/null

# Build and start all containers
echo "Building and starting all containers..."
docker-compose up --build -d

# Check if containers are running
echo "Checking container status..."
sleep 5

if [ "$(docker ps -q -f name=angular-frontend)" ] && \
   [ "$(docker ps -q -f name=code-assistant)" ] && \
   [ "$(docker ps -q -f name=ai-service)" ]; then
  echo "All containers started successfully!"
  echo "====================================================="
  echo "Angular Frontend:  http://localhost:4280"
  echo "Code Assistant:    http://localhost:4281"
  echo "AI Service:        http://localhost:4282"
  echo "====================================================="
else
  echo "Error: Some containers failed to start. Check logs with 'docker-compose logs'"
fi

# Show logs from all containers
echo "Would you like to view container logs? (y/n)"
read -r answer
if [[ "$answer" =~ ^[Yy]$ ]]; then
  docker-compose logs -f
fi