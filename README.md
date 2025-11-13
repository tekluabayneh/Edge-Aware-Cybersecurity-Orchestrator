🛡️ Edge-Aware Cybersecurity Orchestrator
Intelligent Edge Threat Detection and Device Orchestration System
📖 Overview

Edge-Aware Cybersecurity Orchestrator is a distributed system designed to detect, analyze, and respond to potential threats in real time across connected devices.

It’s built to represent a modern edge-first cybersecurity architecture — where devices themselves participate in detecting anomalies and securely communicating with a central orchestrator.

The project is composed of three core components that work together seamlessly:

🟦 Go Backend (Orchestrator) – The control center that manages registered devices, distributes commands, and stores threat reports.

🟩 Go Agent – A lightweight service that runs on each edge device, collects local metrics (CPU, network activity, errors), and executes secure backend commands.

🐍 Python Analyzer – An intelligent analysis service that uses rule-based or machine learning models to detect anomalies and generate alerts based on incoming telemetry.

🟨 TypeScript Client (Dashboard) – A web-based interface for users to monitor devices, review alerts, and visualize network health in real time.

🧠 Problem It Solves

As the number of smart devices and distributed systems continues to grow, traditional centralized security models are no longer enough.

Many organizations struggle with real-time visibility across their networked devices.

Security threats often remain undetected until too late, due to slow data aggregation.

Edge devices lack coordination when reacting to incidents.

This project solves these issues by building a self-coordinating, edge-aware cybersecurity platform, capable of:

✅ Collecting local telemetry in real time
✅ Analyzing anomalies using intelligent rules or models
✅ Sending commands back to devices to respond instantly
✅ Providing operators with an interactive dashboard to monitor everything

⚙️ Architecture Overview
┌────────────────────────┐
│ TypeScript │
│ Client │
│ (Dashboard UI) │
└──────────▲──────────────┘
│ REST / WebSocket
│
┌──────────┴──────────────┐
│ Go Backend │
│ (Orchestrator / API) │
└──────────▲──────────────┘
│ REST / gRPC
│
┌──────────┴──────────────┐
│ Python Analyzer │
│ (AI / Anomaly Engine) │
└──────────▲──────────────┘
│ REST / gRPC
│
┌──────────┴──────────────┐
│ Go Agent │
│ (Edge Device Service) │
└──────────────────────────┘

🧩 Technology Stack
Layer Language Description
Backend / Orchestrator Go Device registration, command dispatch, API, and alert management
Agent (Edge) Go Local telemetry collection and device control
Analyzer Python Anomaly detection and data analysis
Client Dashboard TypeScript (React) Real-time visualization and user control
Communication REST / gRPC Service-to-service interaction
Data Streaming Goroutines / Async Real-time telemetry flow without external broker

🗂️ Project Structure
edge-aware-cybersecurity/
├── backend/ # Go orchestrator
│ ├── cmd/server/
│ ├── internal/
│ └── pkg/
├── agent/ # Go device agent
│ ├── cmd/agent/
│ └── internal/
├── analyzer/ # Python anomaly detection service
│ ├── app/
│ └── tests/
├── client/ # TypeScript dashboard (React)
└── docker-compose.yml # Unified service runner

🚀 Getting Started
Prerequisites

Go 1.22+

Python 3.11+

Node.js 20+

Docker & Docker Compose

Setup

# Clone the repository

git clone ......
cd edge-aware-cybersecurity

# Run all services with Docker

docker-compose up --build

Each service runs in its own container:

Backend → http://localhost:8080

Analyzer → http://localhost:8000

Dashboard → http://localhost:5173

🔒 Security Design

All communication between services is signed and encrypted.

Devices are authenticated when registering with the backend.

Agents can only execute whitelisted commands.

Alerts include severity, origin, and timestamp.

🧠 Future Enhancements

Integrate Kafka or NATS for high-throughput event streaming.

Add machine learning anomaly detection with TensorFlow or PyTorch.

Implement user roles and access control in the dashboard.

Expand agent support for IoT and embedded systems.

🧑‍💻 Author

Teklu Abayneh
Cybersecurity Enthusiast | Backend Developer | Edge Systems Engineer

📜 License

This project is open-source and available under the MIT License.
