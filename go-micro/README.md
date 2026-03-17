## Church Backend API

This repository contains the Go-based backend for the church management system. It uses **Echo v5** for the web framework and **PostgreSQL** for data persistence.

**NOTE:** to run and test server, create a .env file in <root>/go-micro with at least the following vars for now:

- `CHURCH_SECRET`

- `JWT_SECRET`

---

## Technical Stack

* **Language**: Go 1.25+
* **Framework**: Echo v5 (with custom `echox` middleware)
* **Database**: PostgreSQL 16+
* **Authentication**: JWT (HS256)
* **Storage**: Currently MemoryStore for OTP caching

---

## API Documentation for Frontend

All protected routes require a `Bearer` token in the `Authorization` header.

### 1. Authentication

| Endpoint | Method | Body | Description |
| --- | --- | --- | --- |
| `/auth/request` | `POST` | `{"phone": "string"}` | triggers an OTP to the provided phone number. |
| `/auth/verify` | `POST` | `{"phone": "string", "code": "string"}` | verifies OTP and returns a JWT valid for 24 hours(changeable). |

### 2. Member Profile

Routes under the `/members` group are protected by JWT.

| Endpoint | Method | Body | Description |
| --- | --- | --- | --- |
| `/members/profile` | `GET` | N/A | returns the authenticated user's profile and membership status |
| `/members/profile` | `PUT` | `{"first_name": "string", "last_name": "string", "email": "string"}` | updates profile. First and Last names are required. Email is optional |
| `/members/search` | `GET` | `URL params e.g: /members/search?q=ernest&page=1&limit=10` | Takes email, phone, or name parts, finds likely users in pages(extremely low latency & RAM usage, fast fetching time) |

---

## Development Workflow

### Requirements
* Go 1.25+
* Docker and Docker Compose (for streamlining)
* Make (optional, better on Linux)

### Commands

* **Start Environment**: `docker compose up -d` or `make up`
* **Rebuild & Start**: `docker compose up -d --build` or `make build`
* **View Logs**: `make logs` or refer to Makefile for actual command
* **Database Access**: `make db-shell` or refer

---

## Roadmap of pending features

### Core Attendance & Activity

* **Service Tracking**: create schema and endpoints for tracking weekly service types (Sunday Service, Mid-week, etc.).
* **(DONE)Role Based Access**: implement middleware to differentiate between `Admin`, `Member`, and `Visitor` for sensitive routes.
* **(DONE)Search Functionality**: implement server-side filtering and pagination for the member directory (admin only).

## Future Notification & Optimization

* **SMS Gateway Integration**: Replace the current log-based OTP with a functional SMS provider (e.g., MailerSend).
* **Caching Layer**: Integrate `echox` caching on the `GET /members/profile` endpoint to reduce database load for frequent lookups.
* **Attendance System**: Implement `POST /attendance/clock-in` using geo-fencing coordinates or QR code validation.