# Giftwise
GiftWise is an AI-powered gift recommendation app. Enter a recipient's 
hobbies, interests, and budget — get personalized gift ideas generated 
by OpenAI. Save lists privately, rate suggestions, and regenerate ideas.

## Tech Stack
| Component | Technology | Description |
| --- | --- | --- |
| Frontend | Typescript + React | SPA & interface |
| API Gateway | Go (Chi) | REST API, auth, orchestration |
| AI Service | Python + FASTAPI | Gift recommendation generation |
| Database | PostgreSQL | Primary data store |
| AI Model | OpenAI gpt-4o-mini | LLM for recommendations |
| Cloud | AWS (EC2,RDS,S3,CloudFront,IAM,VPC) | Hosting & Infrastructure |
| CI/CD | Github Actions | Build & Deployment Pipeline |

## Prerequisites
- Docker Desktop
- Go 1.25+
- Python 3.14
- golang-migrate (`brew install golang-migrate`)

## Local Setup
1. Clone the repo
2. Copy `.env.example` to `.env` and fill in your values
3. Run `docker compose up --build`

## Go Migrate
First, install go-migrate:
```bash
brew install golang-migrate
```

Verify installation:
```bash
migrate --version
```

Next, run the 000001_init_schema.up.sql file:
```bash
migrate -path api/migrations -database "postgres://your_user:your_password@localhost:5432/your_db?sslmode=disable" up
```

> Use psql to verify the tables were created.

In case of mistake or just want to rollback run:
```bash
# Roll back migrations
migrate -path api/migrations -database "postgres://..." down
```

## Architecture
See [architecture diagram](docs/architecture.html) and [ERD](docs/erd.mermaid).