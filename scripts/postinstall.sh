#!/bin/bash

 Create the system-wide data folder
mkdir -p /var/lib/agent-orchestrator/register

# This ensures your Go app can write the token without sudo later
chmod 777 /var/lib/agent-orchestrator/register

echo "Agent Orchestrator: System directory and permissions initialized."
