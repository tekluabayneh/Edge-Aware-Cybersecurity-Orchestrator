#!/bin/bash

# Use current folder as project root
PROJECT=$(basename "$PWD")

# Create main files
touch "README.md"
touch ".gitignore"
touch "docker-compose.yml"

# Go agent service
mkdir -p "go-agent/config"
touch "go-agent/main.go"
touch "go-agent/go.mod"

# Python analyzer service
mkdir -p "python-analyzer/utils"
touch "python-analyzer/app.py"
touch "python-analyzer/requirements.txt"

# TypeScript backend service
mkdir -p "typescript-backend/src/controllers"
mkdir -p "typescript-backend/src/services"
mkdir -p "typescript-backend/src/models"
mkdir -p "typescript-backend/src/routes"
touch "typescript-backend/package.json"
touch "typescript-backend/tsconfig.json"

# Docs folder
mkdir -p "docs"

echo "Folder structure for $PROJECT created successfully!"

