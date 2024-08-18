# Dating App Service

## Overview

This repository contains a Go-based backend service for a dating app. It is containerized using Docker and includes a MySQL database. The service provides RESTful APIs to handle user interactions, relationships, and other core functionalities.

## Project Structure
```
.
├── api
│   ├── handler
│   │   ├── order_handler.go
│   │   └── user_handler.go
│   ├── presenter
│   │   ├── common.go
│   │   ├── order.go
│   │   └── user.go
│   └── router
│       ├── order_router.go
│       └── user_router.go
├── config
│   └── config.go
├── doc
│   └── thunder-collection_rizkiwhy-dating-app.json
├── middleware
│   └── auth.go
├── pkg
│   ├── order
│   │   ├── entity.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── premium_package
│   ├── relationship_type
│   ├── swipe_history
│   └── user
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Getting Started

1. Install Go: Download and install the latest version of Go from https://golang.org/dl/.

2. Install MySQL: Follow the instructions for your operating system to install MySQL. Create a new database and tables, the query migration is available on `database/migrations.sql`
3. Clone the repository:

```bash
git clone https://github.com/your-username/dating-app-service.git
cd dating-app-service
go run main.go
```