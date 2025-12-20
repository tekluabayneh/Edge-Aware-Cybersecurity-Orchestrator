#!/bin/bash

echo "Installing agent..."

# detect platform
OS=$(uname | tr '[:upper:]' '[:lower:]')

if [[ "$OS" == "linux" ]]; then
    ./agent
elif [[ "$OS" == "darwin" ]]; then
    ./agent
else
    echo "Unsupported OS: $OS"
    exit 1
fi
