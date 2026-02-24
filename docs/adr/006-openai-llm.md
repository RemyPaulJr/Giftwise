# ADR 006 — Use OpenAI API for LLM Integration

**Status:** Accepted  
**Date:** 2026-02-19

---

## Context

GiftWise's core feature is AI-generated gift recommendations. We need an LLM capable of understanding natural language input (hobbies, interests, budget) and returning structured, useful gift suggestions. The model must be accessible via API and reliable enough for a portfolio project.

---

## Decision

We will use the **OpenAI API** (`gpt-4o-mini` as default model) via the official **OpenAI Python SDK** within the Python AI service.

Responses will be requested in structured JSON format using OpenAI's JSON mode or function calling to ensure consistent, parseable output.

---

## Reasons

- OpenAI is the most recognized LLM provider — appearing on a resume carries clear signal
- `gpt-4o-mini` offers a strong balance of capability and cost — suitable for a portfolio project with limited budget
- The Python SDK is mature and well-documented
- JSON mode / function calling ensures the AI returns structured data the Go service can reliably parse
- OpenAI's API is stable and widely used in production systems

---

## Alternatives Considered

**Anthropic Claude API** — A strong alternative with comparable capabilities. Rejected in favor of OpenAI for broader name recognition on resumes, though the Python SDK patterns are nearly identical.

**Google Gemini API** — Considered. Rejected due to less established ecosystem tooling compared to OpenAI at the time of this decision.

**Self-hosted model (Ollama / HuggingFace)** — Rejected. Adds significant infrastructure complexity and cost for marginal learning benefit at this stage. Can be explored post-MVP.

**LangChain** — Considered as an abstraction layer. Rejected for MVP to avoid adding complexity before understanding the underlying OpenAI SDK. Can be added later if prompt chains become complex.

---

## Consequences

- An OpenAI API key is required — must be stored as an environment variable (`OPENAI_API_KEY`) and never committed to source control
- API costs will be incurred per request — `gpt-4o-mini` pricing is low but should be monitored
- The Python service must handle OpenAI API errors gracefully and return meaningful error responses to the Go service
- Prompt engineering will be required to produce consistently useful, structured gift suggestions
- The initial prompt template should be versioned in the codebase so changes are trackable

---

## Example Expected Output Structure

```json
{
  "suggestions": [
    {
      "title": "Hiking Pole Set",
      "description": "Lightweight collapsible poles suitable for trail hiking.",
      "reason": "Based on their interest in outdoor activities and hiking."
    }
  ]
}
```
