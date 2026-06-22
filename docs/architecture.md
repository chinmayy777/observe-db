# ObserveDB Architecture

## What is ObserveDB?

ObserveDB is a learning-focused database implementation written in Go.

## Problem It Solves

Provides a way to store and retrieve data through a simple database interface.

## Version 1 Scope

- Start server
- Accept requests
- Store records in memory
- Retrieve records

## Not Included Yet

- WAL
- Replication
- Clustering
- Transactions
- Distributed consensus

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

The storage package contains business logic.

External clients cannot call Go functions directly.

The HTTP layer exposes storage functionality through endpoints.

Example:

Client -> HTTP API -> Storage Layer

## Responsibilities of API Layer

- Accept requests
- Validate input
- Call storage functions
- Return responses

## What API Layer Should Not Do

- Store data directly
- Contain storage logic
- Manage persistence
- Replace the storage package