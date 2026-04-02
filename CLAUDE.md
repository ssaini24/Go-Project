# CLAUDE.md

This file provides guidance to Claude Code when working with this repository.

## Project Overview

A Go REST API built with the Gin framework. Serves as a test bed for Prism — the AI-powered PR reviewer that analyzes SQL queries and general code quality in pull request diffs.

## Stack

- **Language:** Go 1.24+
- **Framework:** Gin (`github.com/gin-gonic/gin`)
- **Logger:** Uber Zap (`go.uber.org/zap`)
- **Module:** `github.com/ssaini24/Go-Project`

## Project Structure

```
Go-Project/
├── cmd/
│   └── api/
│       └── main.go              # Entry point — initialises config, logger, server
├── internal/
│   ├── config/config.go         # App config loading
│   ├── logger/logger.go         # Zap logger setup
│   ├── middleware/
│   │   └── authorization.go     # Auth middleware (checks Authorization header)
│   ├── router/router.go         # Route registration
│   ├── server/server.go         # Gin server init, runs on :8081
│   ├── user/
│   │   ├── model.go             # User struct
│   │   ├── repository.go        # DB queries — raw SQL via database/sql
│   │   ├── service.go           # Business logic
│   │   └── handler.go           # Gin HTTP handlers
│   └── db/
│       └── queries.sql          # Raw SQL test file (used for Prism SQL rule testing)
└── api/
    └── api.go
```

## Running Locally

```bash
go mod tidy
go run cmd/api/main.go
```

Server starts on `:8081`. All routes are under `/api/v1` and require an `Authorization` header.

## Key Endpoints

| Method | Path | Description |
|---|---|---|
| `GET` | `/api/v1/healthcheck` | Liveness check |
| `GET` | `/api/v1/users/:id` | Get user by ID |
| `GET` | `/api/v1/users` | List all users |
| `GET` | `/api/v1/users/search?name=` | Search users by name |

## Prism Integration

This repo has a GitHub webhook configured pointing to a locally running Prism instance (via ngrok). When a PR is opened or updated, Prism automatically:
1. Fetches the diff
2. Runs SQL static analysis on any SQL queries in the diff
3. Runs code review on Go files via the Code Review Agent
4. Posts inline comments + a summary back to the PR

To trigger a Prism review, open a PR from any feature branch into `main`.
