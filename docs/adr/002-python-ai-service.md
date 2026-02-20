# ADR 002 — Use Python for the AI Recommendation Service

**Status:** Accepted  
**Date:** 2026-02-19

---

## Context

GiftWise needs a service that takes structured input about a gift recipient (hobbies, interests, budget, etc.) and returns AI-generated gift recommendations. This service needs to integrate with an external LLM API and may evolve to include prompt engineering, response parsing, and potentially fine-tuning or retrieval-augmented generation (RAG) in the future.

---

## Decision

We will use **Python** with **FastAPI** as the AI recommendation microservice.

It will expose a single internal HTTP endpoint consumed only by the Go API Gateway. It will use the **OpenAI Python SDK** to communicate with the OpenAI API.

---

## Reasons

- Python is the dominant language in the AI/ML ecosystem — using it here is the most natural fit
- Python is an explicit target resume skill for this project
- FastAPI is modern, async-capable, and auto-generates OpenAPI docs — useful for debugging the service during development
- The OpenAI Python SDK is mature, well-documented, and actively maintained
- Keeping AI logic in Python means the service can later be extended with libraries like LangChain, LlamaIndex, or HuggingFace without touching Go code

---

## Alternatives Considered

**Go with an OpenAI HTTP client** — Rejected. While technically feasible, Go's AI/ML ecosystem is significantly thinner than Python's. This would also eliminate Python from the stack entirely.

**Node.js with OpenAI SDK** — Rejected. TypeScript is already used on the frontend. Python is the more appropriate and resume-relevant choice for an AI service.

**Flask** — Considered as an alternative to FastAPI. Rejected because FastAPI's async support and automatic validation via Pydantic are better suited to a service that may need to handle prompt engineering complexity.

---

## Consequences

- The Python service runs as a separate process/container alongside the Go service
- It is not exposed publicly — only the Go API Gateway calls it internally
- Input/output contracts between Go and Python must be well-defined (JSON over HTTP)
- A separate `requirements.txt` or `pyproject.toml` will manage Python dependencies
- Environment variables (e.g. `OPENAI_API_KEY`) must be securely injected at runtime and never committed to source control
