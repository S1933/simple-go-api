# Simple Go API

A lightweight REST API built with Go's standard library for managing client profiles.

## Project Structure

```
simple-go-api/
├── main.go         # Application entry point and server setup
├── handlers.go     # HTTP request handlers and business logic
├── database.go     # Data models and in-memory database
└── README.md       # This file
```

## Features

- **RESTful API** with support for GET, PATCH, and DELETE operations
- **In-memory database** using Go maps
- **JSON responses** for all operations
- **Error handling** with appropriate HTTP status codes
- **Partial updates** for PATCH operations

## API Endpoints

### Base URL
```
http://localhost:8080
```

### Get Client Profile
Retrieve a client's profile information.

```bash
GET /user/profile?clientId={id}
```

**Example:**
```bash
curl "http://localhost:8080/user/profile?clientId=user1"
```

**Response (200 OK):**
```json
{
  "Email": "email1@gmail.com",
  "Id": "user1",
  "Name": "User One"
}
```

**Error Responses:**
- `403 Forbidden` - Client not found

---

### Update Client Profile
Update a client's name and/or email. Supports partial updates.

```bash
PATCH /user/profile?clientId={id}
Content-Type: application/json
```

**Example:**
```bash
curl -X PATCH "http://localhost:8080/user/profile?clientId=user2" \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Name", "email": "newemail@example.com"}'
```

**Response (200 OK):**
```json
{
  "Email": "newemail@example.com",
  "Id": "user2",
  "Name": "Updated Name",
  "Token": "456"
}
```

**Error Responses:**
- `400 Bad Request` - Missing clientId or invalid JSON
- `404 Not Found` - Client not found

---

### Delete Client Profile
Remove a client profile from the database.

```bash
DELETE /user/profile?clientId={id}
```

**Example:**
```bash
curl -X DELETE "http://localhost:8080/user/profile?clientId=user1"
```

**Response:**
- `204 No Content` - Successfully deleted

**Error Responses:**
- `400 Bad Request` - Missing clientId
- `404 Not Found` - Client not found

---

## Getting Started

### Prerequisites
- Go 1.16 or higher

### Installation & Running

1. **Clone or navigate to the project directory:**
   ```bash
   cd /path/to/simple-go-api
   ```

2. **Build the application:**
   ```bash
   go build *.go
   ```
   This creates a binary named `database`.

3. **Run the server:**
   ```bash
   ./database
   ```

   Or run directly without building:
   ```bash
   go run *.go
   ```

4. **The server will start on port 8080:**
   ```
   Server starting on :8080
   ```

### Initial Data

The application comes pre-populated with two client profiles:

- **user1**
  - Email: email1@gmail.com
  - Name: User One

- **user2**
  - Email: email2@gmail.com
  - Name: User Two

## Architecture

### File Breakdown

#### `main.go`
- Application entry point
- Configures HTTP server on port 8080
- Routes `/user/profile` to `handleClientProfile`

#### `handlers.go`
- **handleClientProfile**: Routes requests to appropriate handlers based on HTTP method
- **GetClientProfile**: Retrieves client data by ID
- **UpdateClientProfile**: Updates client name/email with partial update support
- **DeleteClientProfile**: Removes client from database

#### `database.go`
- **ClientProfile**: Struct defining the data model (Email, Id, Name, Token)
- **database**: In-memory map storing client profiles

### Design Patterns

- **Method-based routing**: Single endpoint handles multiple HTTP methods
- **In-memory storage**: Simple map-based database (data resets on restart)
- **Partial updates**: PATCH only updates fields provided in request body
- **Consistent error handling**: Appropriate HTTP status codes for different error cases

## Development Notes

### Build Commands
```bash
# Build all files together
go build *.go

# Run without building
go run *.go

# Run specific file (won't work due to dependencies)
go build main.go  # ❌ Will fail with "undefined: handleClientProfile"
```

### Important: Multi-file Compilation
When building, you must compile all `.go` files together since `main.go` depends on functions from `handlers.go`, which in turn uses types from `database.go`.

### Stopping the Server
```bash
# Find and kill the process
pkill database

# Or use Ctrl+C in the terminal running the server
```

## Limitations

- **No persistence**: Data is stored in memory and lost on restart
- **No authentication**: Token field exists but isn't validated
- **No input validation**: Limited validation on email format, name length, etc.
- **Single endpoint**: All operations go through `/user/profile`
- **No concurrency protection**: Map access isn't thread-safe for concurrent writes

## Future Enhancements

- Add database persistence (PostgreSQL, SQLite)
- Implement authentication/authorization using Token field
- Add input validation and sanitization
- Create separate routes for each operation
- Add logging and metrics
- Implement proper concurrency handling with mutexes
- Add unit tests
- Add pagination for listing all clients

## License

This is a simple example project for learning purposes.
