# LB-ROCK Load Balancer

![Go Version](https://img.shields.io/badge/Go-1.24%20%7C%201.25%20%7C%201.26-blue)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20MacOS-lightgrey)

A high-performance, concurrent Load Balancer written in Go with a real-time React-based Admin Dashboard.

![System Architecture](./LB-Rock%20system%20arch%20.png)

## Overview

LB-ROCK is a layer 7 load balancer designed for scalability and monitoring. It supports dynamic backend management, health checking, and live traffic metrics visualization.

### Key Features
- **Dynamic Load Balancing**: Distributes traffic across multiple backend servers.
- **Admin Dashboard**: React-based UI for real-time monitoring and node management.
- **REST API Control**: Configure and manage the load balancer state via a dedicated Admin API.
- **Health Checks**: Automated monitoring of backend node availability.
- **Config Sync**: Supports reading and real-time syncing of YAML configurations using **Viper**.

---

## Getting Started

### Prerequisites
- **Go**: 1.24, 1.25, or 1.26
- **Node.js & npm**: (For Dashboard)
- **Docker & Docker Compose**: (Optional)

### Installation
```bash
git clone https://github.com/your-repo/Go-loadBalancer.git
cd Go-loadBalancer
```

### Build Instructions

#### Windows
On Windows, you can use the provided `Makefile`:
```powershell
make build
```

#### macOS / Linux
On Unix-based systems, use the standard Go build command:
```bash
go build -v ./...
```

### Running with Docker (Recommended)
This will spin up the load balancer, the admin dashboard, and three sample backend servers.
```bash
docker compose up --build
```

### Manual Execution
1. **Start the Load Balancer**:
   ```bash
   go run cmd/CLI/main.go start -c config.yaml -p 5000
   ```
2. **Start the Admin Dashboard**:
   ```bash
   cd cmd/AdminApi/adminDashBoard
   npm install
   npm run dev
   ```

---

## Admin API (REST)

The Admin API allows for programmatically managing the load balancer state. Default port: `8086`.

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/status` | GET | Returns the current health and traffic metrics of all nodes. |
| `/add` | POST | Adds a new backend node to the pool. |
| `/remove` | DELETE | Removes a backend node from the pool. |
| `/list` | GET | Lists all configured backend servers. |
| `/check` | GET/POST | Manually triggers a health check on backends. |
| `/logger` | GET | SSE stream for real-time system logs. |
| `/abort` | POST | Triggers a graceful shutdown of the load balancer. |

---

## CLI Commands

Included is a CLI tool for direct management:
- `main.go start`: Launch the balancer.
- `main.go add`: Add a backend.
- `main.go delete`: Remove a backend.
- `main.go status`: View current server status.
- `main.go list`: List all servers in the pool.
- `main.go check`: Run health checks.
- `main.go abort`: Stop the load balancer.

---

## Project Structure
- `cmd/`: Entry points for the Load Balancer, CLI, and Admin API.
- `internal/`: Core logic including load balancing algorithms and health checks.
- `config.yaml`: Global configuration (Sync handled by Viper).
- `servers.yaml`: Initial backend pool configuration.
