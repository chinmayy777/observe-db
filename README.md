# ObserveDB

## Description

A lightweight database written in Go for learning database internals and distributed systems concepts.

## Current Features

- In-memory key-value store
- Put operation
- Get operation
- Delete operation
- Unit tests

## Project Structure

cmd/server - application entry point
internal/storage - storage implementation
docs - architecture and notes

## Running
```bash
go run ./cmd/server
```

## Running Tests

```bash
go test ./...
```

## Future Features

- Persistent storage
- Query support
- WAL
- Replication
- Distributed architecture