# Go WebSocket Server

This is a real-time communication server built with the Go Programming Language. It employs the Model-View-Controller-Service (MVCS) design pattern and a hub-based architecture to manage WebSocket connections and broadcast messages.

## Key Features

- A robust WebSocket server delivering seamless real-time communication between users
- Integrated PostgreSQL database with migrations for reliable data management
- Docker Compose for straightforward development environment setup
- A simple and intuitive API for user management
- A well-structured codebase that is easy to understand and extend
- A Makefile for common development tasks
- Documentation for all endpoints and actions


## Project Structure

The project layout is structured as follows:

- `server`: This is the primary package for the Go WebSocket server code.
- `cmd`: The main application entry point
- `db`: This package contains everything related to database code and migrations
- `dockerfiles`: Contains Dockerfile for setting up your development environment 
- `internal`: This contains several internal packages that follow the MVCS design
  - `user`: Dedicated to user management
  - `ws`: Manages WebSocket communication in line with the hub architecture
- `Makefile`: Useful for common development tasks
- `router`: Houses the HTTP router and API routes
- `util`: A collection of utility functions


## Overview of the MVCS Design Pattern

The MVCS pattern is chosen over traditional MVC due to the additional service layer it introduces. By adding a service layer, business logic and high-level operations are better managed. This separation of concerns improves scalability, reusability, testability, and flexibility. The implementation details include:

- **Model**: The `internal` package has data models and database-related code.
- **View**: The `internal/ws` package handles WebSocket connections and messages.
- **Controller**: The WebSocket handlers manage flow, routing, and interaction with models and the WebSocket hub.
- **Service**: Encapsulates complex business logic and service operations.
- **Repository**: Manages database operations and queries.


![image](https://github.com/Droxt1/WebSOCKET/assets/80992251/95c6060c-fc34-4968-b969-687b5a05676a)

## Essence of the Hub-Based Architecture

- **Hub**: Defined in `internal/ws/hub.go`, it manages WebSocket connections, client additions, removals, and message broadcasts.
- **Client**: Each WebSocket client is represented by a "client" object defined in `internal/ws/client.go`.
- **WebSocket Handlers**: Handlers in `internal/ws/ws_handler.go` interact with the hub to handle WebSocket messages.
- **Initialization and Management**: Hub manages client connections and handles client additions and disconnections.
- **Broadcasting**: Efficient broadcasting of messages to all connected clients.

![image](https://github.com/Droxt1/WebSOCKET/assets/80992251/68d9d97a-506c-40a3-ac5c-74fab5162246)


## Database Schema

This project utilises a "users" table. The SQL code to create the table and associate indexes features below:

```sql
-- Create Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

-- Create Indexes
CREATE INDEX idx_username ON users (username);
CREATE INDEX idx_email ON users (email);
```

## Prerequisites

Ensure you have the following installed on your machine:

- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) 
- [Go](https://golang.org/dl/) is optional, only if you wish to build locally.

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/Droxt1/WebSOCKET.git
   ```

2. Navigate into the project directory:

   ```bash
   cd WebSOCKET
   ```

3. Use Docker Compose to start the development environment:

   ```bash
   make dockerup
   ```

## Working with Database Migrations

Instructions on creating and applying database migrations:

- To create a new migration:

   ```bash
   make createmigration
   ```

- To apply pending migrations:

   ```bash
   make migrateup
   ```
   Ensure to put the SQL code in the `sql.up` file before applying migrations.  

[Documentation](server/documentation.md)
