# Concurrent Job Processing System

A production-inspired concurrent job processing system built in Go, designed to explore core backend and systems engineering concepts such as worker pools, queues, retries, graceful shutdown, dependency injection, and background job execution. **This project does not use any external library or packages, except for .env and live reload**

## Features

* Concurrent worker pool for parallel job execution
* In-memory job queue using Go channels
* Thread-safe in-memory job store
* Executor registry with pluggable job handlers
* Retry support for failed jobs
* Graceful application shutdown
* Structured logging using slog
* REST API for job management
* Dependency injection throughout the application
* Interface-driven architecture for extensibility

---

## Architecture

```text
HTTP API
    ↓
Job Store
    ↓
Job Queue
    ↓
Worker Pool
    ↓
Worker
    ↓
Executor Registry
    ↓
Executors
```

### Components

#### HTTP Layer

Handles job creation, retrieval, and deletion requests.

#### Store Layer

Responsible for job persistence and retrieval.

Current implementation:

* MemoryStore

Future implementations:

* PostgreSQL
* Redis

#### Queue Layer

Responsible for scheduling work for workers.

Current implementation:

* Buffered channel based memory queue

#### Worker Pool

Manages a fixed number of workers responsible for processing jobs concurrently.

#### Executor Registry

Maps job types to their corresponding executors.

#### Executors

Contain the business logic for each job type.

---

## Job Lifecycle

```text
Pending
    ↓
Queued
    ↓
Running
    ├── Completed
    ├── Failed
    └── Retrying
```

Supported statuses:

* pending
* queued
* running
* retrying
* completed
* failed
* cancelled

---

## Supported Job Types

| Job Type           | Description                    |
| ------------------ | ------------------------------ |
| send_email         | Simulates email delivery       |
| send_notification  | Simulates user notifications   |
| generate_thumbnail | Simulates thumbnail generation |
| compress_files     | Simulates file compression     |
| export_user_data   | Simulates user data export     |

---

## API Endpoints

### Health Check

```http
GET /health
```

---

### Create Job

```http
POST /job/create
```

Example request:

```json
{
  "type": "send_email",
  "payload": {},
  "maxRetries": 3
}
```

Example response:

```json
{
  "id": "c7a0cd70-6d1a-4231-9a5a-07a50997a7cf",
  "status": "queued"
}
```

---

### Get All Jobs

```http
GET /jobs
```

Example response:

```json
{
  "total": 5,
  "jobs": []
}
```

---

### Get Job By ID

```http
GET /job/{id}
```

---

### Delete Job

```http
DELETE /job/{id}
```

---

## Graceful Shutdown

The application performs graceful shutdown in the following order:

```text
Stop accepting HTTP requests
        ↓
Close the queue for new jobs
        ↓
Workers finish remaining queued jobs
        ↓
Workers exit
        ↓
Application terminates
```

This ensures no queued jobs are lost during normal shutdown.

---

## Project Structure

```text
.
├── cmd
│   └── server
├── internal
│   ├── api
│   ├── config
│   ├── executor
│   ├── handlers
│   ├── helpers
│   ├── jobs
│   ├── logger
│   ├── queue
│   ├── store
│   └── worker
├── Makefile
└── README.md
```

---

## Running the Project

### Clone Repository

```bash
git clone <repository-url>
cd concurrent-job-processing-system
```

### Install Dependencies

```bash
go mod tidy
```

### Configure Environment

Create a `.env` file:

```env
PORT=8000
LOG_LEVEL=debug
```

### Start Application

```bash
make run
```

Server starts on:

```text
http://localhost:8000
```

---

## Technologies Used

* Go
* net/http
* slog
* sync.WaitGroup
* sync.RWMutex
* Goroutines
* Channels

---

## Future Improvements

* PostgreSQL job store
* Dispatcher for durable jobs
* Dead letter queue
* Scheduled jobs
* Delayed job execution
* Priority queues
* Metrics and monitoring
* Context based cancellation
* Distributed workers
* Job persistence across restarts

---

## Learning Objectives

This project was built to explore:

* Goroutines
* Channels
* Worker pools
* Graceful shutdown
* Interface driven design
* Dependency injection
* Concurrency patterns
* Retry strategies
* Background job processing
* Backend architecture patterns
