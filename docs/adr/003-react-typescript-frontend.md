# ADR 003 — Use React + TypeScript for the Frontend

**Status:** Accepted  
**Date:** 2026-02-19

---

## Context

GiftWise needs a user-facing interface where users can create accounts, log in, enter recipient details, view AI-generated gift recommendations, save lists, and submit ratings and comments. The frontend must be maintainable, typed, and built with tools that are relevant to the current job market.

---

## Decision

We will use **React** with **TypeScript**, bundled with **Vite**.

---

## Reasons

- TypeScript and JavaScript are explicit target resume skills
- React is the most widely used frontend library and appears on a large proportion of full stack job postings
- TypeScript adds static typing to JavaScript, reducing bugs and making the codebase easier to reason about — a strong signal of engineering maturity on a resume
- Vite is the current standard for React project scaffolding — fast, modern, and simple to configure
- React's component model maps naturally to the GiftWise UI (recipient card, gift suggestion card, feedback form, etc.)

---

## Alternatives Considered

**Next.js** — Considered for its SSR capabilities and file-based routing. Rejected for MVP because the added complexity of SSR is not necessary for an authenticated SPA, and it would obscure learning core React patterns.

**Vue or Svelte** — Rejected. React has broader market demand and is more likely to appear on job descriptions targeting the roles GiftWise is being built for.

**Plain JavaScript (no TypeScript)** — Rejected. TypeScript is a target skill and using it from the start reinforces good habits. The type safety also helps when consuming the Go API's JSON responses.

---

## Consequences

- The frontend will be a Single Page Application (SPA) that communicates with the Go API via REST
- All API calls will be made client-side using `fetch` or a lightweight library like `axios`
- Auth tokens (JWT) will be stored in memory or `httpOnly` cookies — this decision will be revisited in the auth ADR
- A component library (e.g. shadcn/ui or Tailwind CSS) may be added for styling — not a core architectural decision
