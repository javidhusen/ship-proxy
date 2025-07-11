# Ship Proxy Project (with Chi)

## Overview

This project contains:
- SHIP Proxy Client (Chi-based)
- Offshore Proxy Server (plain TCP)
- Dockerized setup

## Setup Instructions

### Build Docker Images

```sh
docker build -t ship-client ./client
docker build -t offshore-server ./server
```

### Create Docker Network

```sh
docker network create shipnet
```

### Run Containers

```sh
docker run -d --name proxy-server --network shipnet offshore-server
docker run -d -p 8080:8080 --name proxy-client --network shipnet ship-client
```

## Test Proxy Behavior

### Linux/macOS:

```sh
curl -x http://localhost:8080 http://httpforever.com/
```

### Windows CMD/PowerShell:

```sh
curl.exe -x http://localhost:8080 http://httpforever.com/
```

### Repeated Calls:

Calling the above curl command multiple times should return consistent results.

---