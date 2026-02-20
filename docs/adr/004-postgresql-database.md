# ADR 004 — Use PostgreSQL as the Primary Database

**Status:** Accepted  
**Date:** 2026-02-19

---

## Context

GiftWise needs persistent storage for users, recipient profiles, generated gift suggestions, saved lists, and feedback (star ratings + comments). The data is relational in nature — users have many lists, lists have many suggestions, suggestions have many ratings.

SQL is an explicit target resume skill.

---

## Decision

We will use **PostgreSQL** as the sole database for GiftWise, hosted on **AWS RDS**.

The Go API Gateway will be the only service with direct database access. The Python AI service will not connect to the database — it receives input from Go and returns output to Go, which handles persistence.

---

## Reasons

- PostgreSQL is the most capable and widely used open source relational database — strong resume signal
- The relational data model (users → lists → suggestions → feedback) is a natural fit for SQL
- AWS RDS makes managed PostgreSQL straightforward to set up and operate
- Keeping DB access in one service (Go) enforces a clean separation of concerns
- PostgreSQL supports JSONB if we ever need to store semi-structured data (e.g. raw LLM responses)

---

## Alternatives Considered

**MySQL** — Rejected in favor of PostgreSQL. PostgreSQL has more advanced features, better standards compliance, and is generally preferred in modern backend stacks.

**MongoDB** — Rejected. The data model is clearly relational. Using a document store here would be a poor fit and would eliminate SQL from the stack.

**SQLite** — Rejected for production. Suitable for local development only. RDS provides managed backups, scaling, and reliability that SQLite cannot.

**DynamoDB** — Rejected. While AWS-native, DynamoDB's NoSQL model is not the right fit for this data, and it would eliminate SQL from the stack.

---

## Consequences

- A database schema must be designed and versioned using migrations (e.g. `golang-migrate`)
- The Go service will use `pgx` as the PostgreSQL driver
- Local development will use a PostgreSQL instance via Docker Compose
- AWS RDS will be used for the production environment
- Connection credentials must be managed via environment variables and AWS Secrets Manager in production

---

## Core Schema (Initial)

```
users           — id, email, password_hash, created_at
recipients      — id, user_id, name, hobbies, interests, budget, created_at
gift_lists      — id, user_id, recipient_id, created_at
suggestions     — id, gift_list_id, title, description, reason, created_at
feedback        — id, suggestion_id, user_id, rating (1-5), comment, created_at
```
