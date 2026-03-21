🛡️ Edge-Aware Cybersecurity Orchestrator
Intelligent Edge Threat Detection and Device Orchestration System
📖 Overview
Edge-Aware Cybersecurity Orchestrator is a distributed system designed to detect, analyze, and respond to potential threats in real time across connected devices. It represents a modern edge-first architecture where devices participate in their own defense.

The Core Components:
🟦 Go Backend (Orchestrator): The command center managing device registration and threat intelligence.

🟩 Go Agent (Edge Service): A lightweight service running on edge devices, collecting telemetry and executing secure commands.

🐍 Python Analyzer: An AI-driven engine using machine learning to flag anomalies in real-time data.

🟨 TypeScript Client (Dashboard): A React-based interface for global network visualization and alert management.

🚀 Installation & Setup
You can now install the Edge Agent using professional installers or run the entire stack via Docker.

1. For End Users (Installers)
If you just want to run the Agent on a device, download the latest version from the Releases Page.

Windows: Download Agent_Setup_amd64.exe. Run the installer to set up the service and automatic permissions.

Linux (Ubuntu/Debian): Download agent_1.0.x_amd64.deb.

Bash
sudo dpkg -i agent_1.0.x_amd64.deb
The Linux installer automatically configures /var/lib/agent-orchestrator/ with the necessary write permissions for telemetry storage.

2. For Developers (Docker Compose)
To run the full ecosystem (Backend, Analyzer, Agent, and Dashboard) at once:

Bash
# Clone the repository
git clone https://github.com/tekluabayneh/Edge-Aware-Cybersecurity-Orchestrator.git
cd Edge-Aware-Cybersecurity-Orchestrator

# Spin up the entire stack
docker-compose up --build
Orchestrator API: http://localhost:8080

AI Analyzer: http://localhost:8000

Dashboard UI: http://localhost:5173

⚙️ Architecture Overview
Code snippet
graph TD
    subgraph Cloud/Central
    A[TypeScript Dashboard] <-->|REST/WS| B[Go Backend]
    B <-->|gRPC/REST| C[Python AI Analyzer]
    end
    
    subgraph Edge Devices
    B <-->|Secure Channel| D[Go Agent 1]
    B <-->|Secure Channel| E[Go Agent 2]
    end
🔒 Security & Edge Design
Zero-Config Permissions: Windows and Linux installers handle filesystem permissions automatically so the Agent can store JWT tokens securely.

Whitelisted Execution: Agents only execute backend-signed commands to prevent remote code execution (RCE).

Data Sovereignty: Telemetry is processed locally before being summarized for the orchestrator.

🗂️ Project Structure
agent/: Go source for the edge service.

backend/: Go source for the orchestrator and database logic.

analyzer/: Python scripts for threat detection logic.

client/: React/Vite dashboard source.

scripts/: Post-install automation for Linux environments.

installer.iss: Inno Setup configuration for Windows production builds.

🧠 Database Schema
Detailed schema relations can be viewed here:
https://drawsql.app/teams/man-21/diagrams/edage-aware-security-db/embed

🧑‍💻 Author
Teklu Abayneh 
Cybersecurity Enthusiast | Full-Stack Engineer| Edge Systems Engineer

<!-- https://drawsql.app/teams/man-21/diagrams/edage-aware-security-db/embed -->

