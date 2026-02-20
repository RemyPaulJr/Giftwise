# ADR 008 — Deploy on AWS

**Status:** Accepted  
**Date:** 2026-02-19

---

## Context

GiftWise needs a production hosting environment. AWS is an explicit target resume skill. We need to choose which AWS services to use and keep the setup achievable for a junior developer working on a portfolio project.

---

## Decision

We will deploy GiftWise on **AWS** using the following services:

| Component | AWS Service |
|---|---|
| Go API Gateway | ECS (Fargate) or EC2 |
| Python AI Service | ECS (Fargate) or EC2 (same VPC) |
| PostgreSQL Database | RDS (PostgreSQL) |
| Secrets (API keys, JWT secret) | AWS Secrets Manager or SSM Parameter Store |
| Container Registry | ECR (Elastic Container Registry) |
| Frontend (SPA) | S3 + CloudFront |
| CI/CD | GitHub Actions → ECR → ECS deploy |

The MVP may start with EC2 for simplicity and migrate to ECS as familiarity grows. Either approach is resume-valid.

---

## Reasons

- AWS is the most widely used cloud provider and the most commonly listed on job descriptions
- ECS Fargate is a managed container platform — no server management required
- RDS provides managed PostgreSQL with automated backups
- S3 + CloudFront is the standard, low-cost pattern for hosting a static SPA
- ECR keeps container images within the AWS ecosystem
- GitHub Actions is free for public repos and pairs naturally with an AWS deployment pipeline

---

## Alternatives Considered

**GCP (Google Cloud)** — Rejected. AWS has broader job market presence for backend roles.

**Heroku** — Rejected. Heroku is simpler but no longer free and carries less resume weight than AWS.

**Render / Railway** — Rejected for the same reason. Useful for quick deploys but not resume-differentiated.

**Kubernetes (EKS)** — Rejected for MVP. EKS adds significant operational complexity that is not appropriate for this stage. ECS Fargate achieves container orchestration with far less overhead.

**Vercel for frontend** — Considered. Rejected to keep the full deployment within AWS for consistency and resume coherence.

---

## Consequences

- An AWS account is required — the free tier covers most MVP usage but costs should be monitored
- IAM roles and policies must be configured carefully — principle of least privilege
- Infrastructure will initially be set up manually via the AWS console, with a note to explore Terraform or CDK post-MVP
- Environment variables and secrets must never be hardcoded — use Secrets Manager or SSM
- A basic architecture diagram should be added to `docs/` once infrastructure is provisioned

---

## MVP Deployment Simplification

For the earliest MVP, it is acceptable to:
- Run both services on a single EC2 instance using Docker Compose
- Use RDS in the same VPC
- Deploy the frontend to S3 + CloudFront

This keeps costs near zero and avoids ECS complexity until the app is working end-to-end.
