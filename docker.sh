#!/bin/bash

echo "Building image"
docker build -t nearest-prime-app .
echo "Running docker"
docker run -p 8080:8080 nearest-prime-app