# ADR 001 — Use Go for the API Gateway

**Status:** Accepted  
**Date:** 2026-02-19

---

## Context

GiftWise needs a backend API layer responsible for handling HTTP requests, authentication, user management, saved gift lists, and feedback storage. This service is the central orchestrator — it talks to the database, the frontend, and the internal Python AI service.

We needed a language that is performant, has strong typing, is well-suited to building REST APIs, and is a high-demand skill for backend engineering roles.

---

## Decision

We will use **Go** as the language for the API Gateway service.

We will use the **Chi** router (`go-chi/chi`) for HTTP routing due to its lightweight footprint and idiomatic Go design. It is stdlib-compatible and does not impose an opinionated framework structure.

---

## Reasons

- Go is explicitly a target resume skill for this project
- Go's standard library and ecosystem make REST API development straightforward
- Go compiles to a single binary, making deployment to AWS simple
- Strong typing and explicit error handling encourage good engineering habits
- High concurrency via goroutines is a natural fit for an API server handling multiple requests
- Chi is minimal and composable — better for learning fundamentals than an opinionated framework like Echo or Fiber

---

## Alternatives Considered

**Node.js (TypeScript)** — Rejected. TypeScript is already used on the frontend. Using it on the backend too would reduce the breadth of the resume and miss the Go requirement entirely.

**Python (FastAPI)** — Rejected. Python is reserved for the AI service. Keeping concerns separated by language reinforces the microservice boundary and keeps both services focused.

**Echo or Fiber (Go frameworks)** — Considered as alternatives to Chi. Rejected in favor of Chi for its closer alignment to the Go standard library, which better supports learning Go idioms.

---

## Consequences

- The team (you) will need to learn Go syntax, error handling patterns, and package management via Go modules
- All database access (PostgreSQL) will be handled through this service using `pgx` or `database/sql`
- This service will expose a REST API consumed by the React frontend
- Inter-service communication to the Python AI service will be done over HTTP (internal)
