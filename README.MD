# Nearshore Test API

## Overview

The Nearshore Test API provides a simple web application for managing devices and their respective firmware versions. This repository contains a Go application which enables Create, Read, Update, and Delete (CRUD) operations on devices and firmware through a set of defined API endpoints.

## 🛠️ Technologies

- **Go 1.21+**: Main programming language.
- **Gin**: Web framework used.
- **Godotenv**: Utilized for handling environment variables.
- **PostgreSQL**: Database used for storing data.
- **Google UUID**: For generating UUIDs.
- **Bcrypt**: For hashing and verifying passwords.
- **Docker** & **Docker Compose**: For containerization and managing multi-container Docker applications, respectively.

## 📁 Project Structure

```
nearshore-test/
│
|── .dockers
|   └── postgres
|       └── Dockerfile
|       └── init.sql
├── internal
│   └── core
│       └── dc-nearshore
│           └── endpoint.go
│       └── device
│           └── endpoint.go
│           └── repository.go
│           └── service.go
|       └── domain
│           └── device.go
│           └── firmware.go        
│       └── firmware
│           └── endpoint.go
│           └── repository.go
│           └── service.go
│   └── middlewares.go
│   └── urls.go
├── pkg/
│   ├── bootstrap
│   │   └── bootstrap.go
│   ├── helper
│   │   └── password.go
│   ├── meta
│   │   └── meta.go
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── README.md
```


## 🚀 Getting Started

### Prerequisites

- Docker & Docker Compose
- Postman or Insomnia (For API testing, optional)

### Setup & Running

1. **Clone the Repository**
    ```sh
    git clone https://github.com/cesarcruzc/nearshore_test
    ```

2. **Configure Environment Variables**
   Copy the `.env.example` file to `.env` and update the variables accordingly.

   ```sh
   cp .env.example .env

3. Run the following command to start the application:
    ```sh
    docker-compose up
    ```

4. The application will be available at `http://localhost:8080`

## 🌐 API Endpoints

### 🔐 Authentication

The API employs Basic Authentication for accessing the **Devices** and **Firmware** endpoints. You must provide a valid username and password to interact with these endpoints.

### Using Basic Auth with cURL

When using cURL to make requests to the API, use the `-u` flag to pass the username and password:

```sh
curl -X GET http://localhost:8080/devices -u 'your-username:your-password'
```

### Devices

- GET /devices - Get all devices
- POST /devices - Create a device
- GET /devices/:id - Get a device
- PUT /devices/:id - Update a device
- DELETE /devices/:id - Delete a device

### Firmwares

- GET /firmwares - Get all firmwares
- POST /firmwares - Create a firmware
- GET /firmwares/:id - Get a firmware
- PUT /firmwares/:id - Update a firmware
- DELETE /firmwares/:id - Delete a firmware

### Other endpoints

- GET / - Service
- GET /health - Healthcheck

## 🔗 API Testing with Insomnia

Import the insomnia.json file into Insomnia to test the endpoints with predefined requests.

## 🧑‍💻 Author

- César Cruz
- cc.cruz.caceres@gmail.com