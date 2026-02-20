# ADR 007 — Microservice Split: Go API + Python AI Service

**Status:** Accepted  
**Date:** 2026-02-19

---

## Context

GiftWise has two distinct backend responsibilities: general API logic (auth, data persistence, routing) and AI-powered recommendation generation. These could be implemented in a single service or split into separate services. The right approach affects the resume story, maintainability, and deployment complexity.

---

## Decision

We will split the backend into **two services**:

1. **Go API Gateway** — handles all client-facing HTTP requests, auth, database access, and orchestration
2. **Python AI Service** — an internal-only service responsible solely for generating gift recommendations via OpenAI

The Go service calls the Python service over HTTP on an internal network. The Python service is never exposed to the public internet.

---

## Communication Contract

```
Go API Gateway  →  POST http://ai-service:8001/recommend
                   Body: { "hobbies": [...], "interests": [...], "budget": "..." }
                   Response: { "suggestions": [...] }
```

---

## Reasons

- Demonstrates understanding of microservice architecture — a strong talking point in interviews
- Allows each service to use the most appropriate language and ecosystem for its responsibility
- The Python service can evolve independently (e.g. swap models, add RAG) without touching Go code
- Clean separation of concerns makes each service easier to reason about and test individually
- Mirrors patterns used in real-world AI product backends

---

## Alternatives Considered

**Monolith (Go only, calling OpenAI directly)** — Rejected. While simpler, it eliminates Python from the stack and reduces the architectural complexity that makes this project resume-worthy.

**Monolith (Python only)** — Rejected. Eliminates Go from the stack.

**gRPC for inter-service communication** — Considered as an alternative to HTTP. Rejected for MVP due to added complexity (proto files, code generation). Plain HTTP/JSON is sufficient and faster to implement.

**Message queue (e.g. SQS) between services** — Considered for async processing. Rejected for MVP — synchronous HTTP keeps the flow simple and easier to debug.

---

## Consequences

- Both services must be running for the app to function — local Docker Compose will manage this
- Network latency between services is introduced but negligible for this use case
- Each service will have its own `Dockerfile`
- A `docker-compose.yml` at the repo root will wire both services together for local development
- In production on AWS, both services will run on the same VPC so the Python service is not publicly reachable
