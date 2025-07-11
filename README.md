# Student API

A simple RESTful API for managing student records, built with Go.

## Project Overview

This project provides endpoints to create and retrieve student information. It demonstrates a clean architecture with separation of concerns, validation, and structured error handling. The API is suitable for learning, prototyping, or as a starting point for more complex systems.

## Features
- Create a new student
- Retrieve student details by ID
- Input validation and error responses

## Folder Structure

```
student_api/
├── cmd/
│   └── students-api/         # Entry point for the application (main.go)
├── internal/
│   ├── config/               # Configuration management
│   ├── http/
│   │   └── handlers/
│   │       └── student/      # HTTP handlers for student endpoints
│   ├── storage/
│   │   ├── database/         # Database connection and logic
│   │   └── storage.go        # Storage interface and implementation
│   ├── types/                # Type definitions (e.g., Student struct)
│   └── utils/
│       └── response/         # Utility functions for HTTP responses
├── go.mod                    # Go module definition
├── go.sum                    # Go module checksums
```

## Prerequisites
- Go 1.18 or higher

## How to Run

1. **Clone the repository:**
   ```sh
   git clone <your-repo-url>
   cd student_api
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Run the API server:**
   ```sh
   go run ./cmd/students-api/main.go
   ```

4. **API Endpoints:**
   - `POST /students` — Create a new student
   - `GET /students/{id}` — Get student by ID

## Example Request

**Create Student:**
```sh
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "age": 20}'
```

**Get Student by ID:**
```sh
curl http://localhost:8080/students/1
```

## License

This project is for educational purposes. 