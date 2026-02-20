# ADR 005 — Use JWT for Authentication

**Status:** Accepted  
**Date:** 2026-02-19

---

## Context

GiftWise requires user accounts. Users must be able to register, log in, and access their own private data (saved lists, feedback). We need an authentication mechanism that is stateless, works well with a REST API, and is manageable without a third-party auth provider for MVP.

---

## Decision

We will implement **JWT (JSON Web Token)** based authentication, managed entirely within the Go API Gateway.

On login, the server issues a signed JWT. The client stores it and sends it in the `Authorization: Bearer <token>` header on subsequent requests. The Go service validates the token on protected routes.

Tokens will be stored client-side in **memory** (React state) for security, with a short expiry (e.g. 15 minutes access token). A refresh token strategy can be added post-MVP.

---

## Reasons

- JWT is stateless — no session store required, which simplifies the backend
- Industry-standard pattern for REST API auth — strong resume signal
- Keeps auth in-house, avoiding third-party dependencies for MVP
- Go has mature JWT libraries (e.g. `golang-jwt/jwt`)
- Storing tokens in memory (not localStorage) mitigates XSS risk

---

## Alternatives Considered

**AWS Cognito** — Considered as a managed auth solution. Rejected for MVP because it abstracts away the implementation details that are valuable to learn. Can be revisited post-MVP.

**Session-based auth (cookies + server-side sessions)** — Rejected. Requires a session store (Redis or DB), which adds infrastructure complexity. JWT keeps the service stateless.

**Auth0 / Clerk** — Rejected for the same reason as Cognito. Third-party auth services are valuable in production but reduce the learning surface area.

**localStorage for token storage** — Rejected due to XSS vulnerability. In-memory storage is safer for MVP.

---

## Consequences

- Go must implement registration (`POST /auth/register`) and login (`POST /auth/login`) endpoints
- A JWT signing secret must be stored as an environment variable and never committed to source control
- Protected routes in Go will use middleware to validate the JWT before allowing access
- The React frontend must handle token expiry gracefully (e.g. redirect to login)
- Password hashing will use **bcrypt** via Go's `golang.org/x/crypto/bcrypt`
- Refresh token flow is out of scope for MVP but noted as a future improvement
