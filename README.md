# DevOps Intern Assignment: Nginx Reverse Proxy + Docker Compose

A Docker-Compose setup that runs:

- **Service 1**: Go HTTP server on port 5001  
- **Service 2**: Python/Flask server on port 5002  
- **Nginx**: Reverse proxy on port 8080 routing `/service1/` → Service 1 and `/service2/` → Service 2

---

## Prerequisites

- Docker & Docker Compose installed  
- Git (to clone this repo)

---

## Quick Start

1. **Clone the repo**  
   ```bash
   git clone https://github.com/kapilsharma082/Assignment.git
   cd Assignment
   
## Build and run all services
docker compose up --build

## TO VERIFY
# Root page
curl http://localhost:8080/

# Service 1 health
curl http://localhost:8080/service1/health

# Service 2 health
curl http://localhost:8080/service2/health
## How Routing Works
Nginx listens on port 80 inside its container (mapped to host 8080).

Requests to /service1/ are proxied to service1:5001.

Requests to /service2/ are proxied to service2:5002.

