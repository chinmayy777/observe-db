# ObserveDB Architecture

## What is ObserveDB?

ObserveDB is a learning-focused database implementation written in Go. The goal of the project is to understand how databases and backend services are built by implementing each component from scratch.

## Problem It Solves

ObserveDB provides a simple key-value database that allows clients to store and retrieve data through an HTTP API.

## Version 1 Scope

- Start HTTP server
- Accept HTTP requests
- Store records in memory
- Retrieve records
- Delete records
- Basic API layer
- Storage package

## Not Included Yet

- WAL
- Persistent storage
- Replication
- Clustering
- Transactions
- Distributed consensus
- Authentication
- Concurrency Protection

## Future API Ideas

Store records using key-value pairs.

Possible operations:

- PUT key/value
- GET by key
- DELETE by key

Possible request format:

{
    "key" : "name"
    "value" : "Chinmay"
}

Possible response format:

{
    "success" : true
}

## Why an HTTP Layer Exists

The storage package contains the core database logic.

External clients cannot call Go functions directly.

The HTTP layer exposes storage functionality through endpoints.

Example:

Client -> HTTP API -> Storage Layer

## Responsibilities of API Layer

- Accepting HTTP requests
- Validating HTTP methods
- Decoding JSON request bodies
- Reading query parameters
- Calling storage functions
- Returning JSON responses
- Returning appropriate HTTP status codes

## What API Layer Should Not Do

- Store data directly
- Contain storage logic
- Manage persistence
- Replace the storage package

## Current API

# Health Check

GET /health

Response: ObserveDB is running

# Store Data

POST /put

Request: 
{
    "key" : "name"
    "value" : "Chinmay"
}
Response:
{
    "status" : "stored"
}

# Retrieve Data

GET /get?key=name

Response:
{
    "value" : "Chinmay"
}

If the key doesn't exist:
404 Not Found

## Data Flow

# Store Request

Client -> POST /put -> PutHandler -> Store.Put() -> Memory

# Retrieve Request

Client -> GET /get -> GetHandler -> Store.Get() -> JSON Response


## Future Roadmap

Upcoming Improvements:

- DELETE endpoint
- API tests
- Better error handling
- Persistent storage
- Write-Ahead Log (WAL)
- Concurrency support
- Configuration management
- Storage engine improvements