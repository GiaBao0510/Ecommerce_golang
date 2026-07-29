# E-Commerce Backend With Golang

This repository is a backend project for an e-commerce platform built with Golang. The project focuses on API design, database modeling, authentication foundations, role and permission management, logging, caching infrastructure, and containerized local development.

The codebase is designed as a learning-oriented but production-minded backend system. It separates HTTP handling, business logic, data access, database queries, middleware, configuration, and shared utilities into clear packages so the project can grow feature by feature without becoming tightly coupled.

## Project Overview

The application provides the foundation for an online commerce system with users, products, brands, categories, orders, carts, coupons, reviews, notifications, roles, permissions, OAuth accounts, refresh tokens, and user activity logs.

Current implemented foundations include:

- Golang backend using the Gin HTTP framework.
- PostgreSQL schema for core e-commerce domains.
- SQL query generation with sqlc.
- Repository, service, controller, mapper, middleware, and router layers.
- Redis connection initialization for future caching and OTP workflows.
- Structured logging with Zap and file rotation with Lumberjack.
- Request tracing through trace ID middleware.
- Standardized API response helpers.
- Health check endpoints.
- Role-based access control database structure.
- Dockerfile and Docker Compose setup for local infrastructure.
- Unit/integration testing structure with Testify and Bruno API collections.

## Main Features

### User And Account Management

The project includes user-related database tables and API layers for creating, retrieving, updating, deleting, and searching users by email or phone number. The user schema also supports email verification, phone verification, avatar URL, status tracking, refresh tokens, OAuth accounts, and role assignment.

### Role And Permission Management

The database model includes roles, permissions, and role-permission mapping. This provides the foundation for role-based access control, such as admin, manager, shop owner, and customer permissions.

### Product And Catalog Model

The database schema supports product catalog management through brands, categories, sub-categories, product details, product images, reviews, and inventory-related fields such as price, quantity, production date, and expiration date.

### Cart, Order, Payment, And Coupon Model

The schema includes shopping carts, cart items, orders, order details, payment methods, user payment methods, and order coupons. This gives the project a strong base for building checkout and order processing flows.

### Logging And Observability Foundation

The project uses Zap for structured application logging and Lumberjack for log rotation. Middleware also adds request-level trace IDs and HTTP request logging, which makes it easier to debug request flows and prepare the system for centralized observability.

### Redis Infrastructure

Redis is initialized as part of the application startup and included in Docker Compose. It is planned to support OTP verification, short-lived authentication data, rate limiting, and caching for high-read endpoints.

### Docker-Based Local Development

The project includes a multi-stage Dockerfile for building the Go application and a Docker Compose setup for PostgreSQL and Redis. This makes local setup more consistent and prepares the project for future deployment workflows.

## Planned Advanced Capabilities

The project is being extended with the following advanced backend capabilities. These are part of the target architecture and roadmap for the next development phases.

### JWT Authentication And Authorization

JWT authentication will be used to secure protected APIs. Access tokens will identify authenticated users, while refresh tokens will support long-lived sessions. Combined with the role and permission tables, the system can enforce role-based authorization for admin, manager, shop owner, and customer actions.

### Google OAuth2 Login

Google OAuth2 login will allow users to authenticate with their Google account. The `OAUTH_ACCOUNT` table already provides a database foundation for linking external provider accounts to internal users.

### OTP Verification With Redis

Redis will be used to store short-lived OTP codes for email or phone verification. This is a good fit because OTP data is temporary, requires fast read/write access, and should expire automatically.

### Loki And Grafana Logging

The logging system will be extended to send application logs to Loki and visualize them through Grafana. This will help monitor request errors, API latency, authentication issues, and background job behavior in a centralized dashboard.

### Database Query Optimization

The schema already includes indexes for important fields such as user email, phone number, product name, order user ID, order status, cart user ID, OAuth provider ID, and review product ID. Future optimization work will include analyzing slow queries, improving pagination, avoiding unnecessary joins, and adding indexes based on real query patterns.

### WebSocket Realtime Chat And Notifications

WebSocket support will be added for realtime features such as customer-shop chat, order status notifications, promotion alerts, and admin/customer notifications. This will improve user experience for workflows that should not rely only on polling.

### Gemini API For Product Description Enhancement

Gemini API integration will be used to improve product descriptions. For example, sellers can provide a short product draft and the system can generate a clearer, more attractive, SEO-friendly product description before publishing.

### Elasticsearch For Product Search

Elasticsearch will be added to improve product search and filtering. It can support full-text search, typo-tolerant queries, category filters, brand filters, price ranges, and ranking based on relevance or popularity.

### Kafka Message Queue

Kafka will be introduced for asynchronous event processing. In an e-commerce system, Kafka is useful for events such as order created, payment completed, inventory updated, user registered, notification requested, and product updated.

Possible Kafka use cases in this project:

- Publish an `order.created` event after a user places an order.
- Trigger notification delivery without blocking the checkout API.
- Sync product data from PostgreSQL to Elasticsearch.
- Process user activity logs asynchronously.
- Update analytics or reporting services from domain events.
- Prepare for future CDC integration with Debezium.

Using Kafka keeps the main API faster and makes the system easier to extend when more services are added.

## Technology Stack

| Area | Technology |
| --- | --- |
| Language | Golang |
| HTTP Framework | Gin |
| Database | PostgreSQL |
| Query Layer | sqlc |
| Cache / Temporary Data | Redis |
| Logging | Zap, Lumberjack |
| Message Queue | Kafka |
| Search Engine | Elasticsearch |
| Observability | Loki, Grafana |
| Authentication | JWT, OAuth2 |
| AI Integration | Gemini API |
| Containerization | Docker, Docker Compose |
| Testing | Go testing, Testify, Bruno |

## Project Structure

```text
cmd/
  server/              Application entry point
  cli/                 CLI and experiment entry points

internal/
  controller/          HTTP request handlers
  service/             Business logic layer
  repository/          Data access interfaces and implementations
  database/            sqlc generated database code
  dto/                 Request and response DTOs
  mapper/              Model and DTO mapping helpers
  middleware/          HTTP middleware
  models/              Domain and request models
  routers/             API route registration
  initialize/          Application initialization
  util/                Internal utility helpers
  wire/                Dependency injection setup

pkg/
  apperrors/           Application error types
  logger/              Logger setup
  loghelper/           Logging helpers
  response/            Standard API response helpers
  setting/             Configuration structures

sql/                   Database schema and sqlc queries
migrations/            Seed and migration-related SQL files
test/                  Tests and API collections
diagrams/              Database design files
```

## Local Development

Start PostgreSQL and Redis:

```bash
docker compose up -d
```

Run the API server:

```bash
make run
```

Or run directly:

```bash
go run ./cmd/server/main.go
```

The application starts on:

```text
http://localhost:8080
```

## Development Direction

This project is still evolving. The next development focus is to complete authentication, authorization, OTP verification, product APIs, order flows, realtime notifications, search integration, message queue processing, and observability.

The long-term goal is to turn this codebase into a realistic e-commerce backend that demonstrates backend engineering practices such as clean architecture layering, secure authentication, optimized database access, asynchronous processing, containerized infrastructure, and operational monitoring.
