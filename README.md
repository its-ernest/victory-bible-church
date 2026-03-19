# Church system and app

## Backend docs:
[go-micro/README.md](go-micro/README.md)

A unified ecosystem for church management, member engagement, and media distribution. This project consists of a high-performance **Go (Echo v5) Backend** and a modern **Frontend/Mobile** interface.



## Project Vision
To bridge the gap between the sanctuary and the digital world by providing members with easy access to sermons, ministry involvement, and profile management.

---

## The Stack

| Component | Technology | Role |
| :--- | :--- | :--- |
| **Backend** | Go 1.25+, Echo v5, PGX | High-concurrency API & Business Logic |
| **Database** | PostgreSQL 16 | Relational data & Sermon indexing |
| **Auth** | JWT (v5) | Secure stateless authentication |
| **Infrastructure** | Docker & Compose | Containerized development & deployment |

---

## Project Structure

```text
.
├── go-micro/               # backend Service (Golang)
│   ├── internal/           # private application code
│   │   ├── models/         # database & API structs
│   │   ├── routes/         # handlers & Service logic (Sermons, Ministries, etc.)
│   │   └── repository/     # postgres logic
│   ├── Dockerfile          # multi-stage Go build
│   ├── docker-compose.yml
│   └── Makefile            # dev automation commands
├── frontend/               # web/Mobile Client (React/Next.js)
├── docker-compose.yml      # orchestrates API and Database
└── init.sql                # database schema and seeds
```

---

## Quick Start (Development)

### 1. Prerequisites
* Docker & Docker Desktop
* Postman (for API testing)
* Make (Optional)

### 2. Launch the Environment
From the root directory, navigate to the backend folder and use the automated commands:

```bash
cd go-micro
make build   # Builds images and starts containers
make logs    # Tails the API output
```

### 3. API Access
The backend is exposed at `http://localhost:8080`.
* **Public**: `GET /sermons`, `GET /ministries`
* **Private**: Requires `Authorization: Bearer <JWT>` for protected routes such as: `GET /profile` and `POST /ministries/join`.

---

## Current Features
* **Member Management**: Profile updates and status tracking.
* **Ministry Hub**: Explore departments (Choir, Media) and join them.
* **Sermon Library**: Filterable media archive with tag support.
* **RBAC**: Role-based access for Admin-only uploads.

---

## Collaboration Workflow
* **Backend**: Managed in the `go-micro` directory. API changes should be documented for the frontend team.
* **Frontend**: Pulls data from the defined Echo routes.
* **Database**: Schema changes must be added to `init.sql` to stay synced across environments.

### Fork and clone your copy of this repo, add improvements, push. Then open a pull request to merge the changes together here
### Include test files to ensure integrity and predictable performance