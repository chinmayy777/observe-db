# Week 2 Review

## What HTTP layer works

Implemented a basic HTTP server using Go's net/http package.

Current endpoints:

- GET /health
- POST /put
- GET /get
- DELETE /delete

The API successfully communicates with the storage package.

Requests are validated before being processed.

JSON request bodies are decoded correctly.

Responses are returned in JSON format.

---

## What still feels rough

Handlers contain repeated validation logic.

There is no router.

Error responses are inconsistent.

No logging exists.

No persistence exists.

Data is lost whenever the server restarts.

---

## Tests that pass

Storage tests:

- Put
- Get
- Delete
- Missing key
- Overwrite value

API tests:

- Valid PUT
- Valid GET
- Valid DELETE
- Missing key
- Bad JSON
- Missing record

All tests pass using:

go test ./...

---

## What next week should focus on

Improve the HTTP layer.

Reduce duplicate code.

Learn better request validation.

Prepare for persistent storage in future weeks.