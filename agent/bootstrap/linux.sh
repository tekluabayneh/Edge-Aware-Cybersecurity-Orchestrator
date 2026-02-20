#!/bin/bash

APP_PATH="$(realpath ../bootstrap/linux.sh)"

SERVICE_PATH="/etc/systemd/system/myagent.service"

echo "Setting up persistence on Linux..."

sudo tee $SERVICE_PATH > /dev/null <<EOF
[Unit]
Description=MyAgent Service
After=network.target

[Service]
ExecStart=$APP_PATH
Restart=always

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl daemon-reload
sudo systemctl enable myagent
sudo systemctl start myagent

echo "Persistence enabled on Linux."
