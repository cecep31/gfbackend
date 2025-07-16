# Users API Documentation

This document describes the users REST API endpoints added to the GoFrame backend.

## Endpoints Overview

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users` | Get all users with pagination |
| GET | `/users/:id` | Get user by ID |
| POST | `/users` | Create new user |
| PUT | `/users/:id` | Update user |
| DELETE | `/users/:id` | Delete user |

## API Details

### 1. Get All Users
- **Method**: `GET`
- **Path**: `/users`
- **Query Parameters**:
  - `page` (optional): Page number (default: 1)
  - `limit` (optional): Items per page (default: 10)
- **Response**:
```json
{
  "users": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "2025-07-16 13:45:00",
      "updated_at": "2025-07-16 13:45:00"
    }
  ],
  "total": 1
}
```

### 2. Get User by ID
- **Method**: `GET`
- **Path**: `/users/:id`
- **Path Parameters**:
  - `id`: User ID (required)
- **Response**:
```json
{
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2025-07-16 13:45:00",
    "updated_at": "2025-07-16 13:45:00"
  }
}
```

### 3. Create New User
- **Method**: `POST`
- **Path**: `/users`
- **Request Body**:
```json
{
  "name": "Alice Johnson",
  "email": "alice@example.com",
  "password": "secure123"
}
```
- **Response**:
```json
{
  "user": {
    "id": 3,
    "name": "Alice Johnson",
    "email": "alice@example.com",
    "created_at": "2025-07-16 13:45:00",
    "updated_at": "2025-07-16 13:45:00"
  }
}
```

### 4. Update User
- **Method**: `PUT`
- **Path**: `/users/:id`
- **Path Parameters**:
  - `id`: User ID (required)
- **Request Body**:
```json
{
  "name": "Updated Name",
  "email": "updated@example.com"
}
```
- **Response**:
```json
{
  "user": {
    "id": 1,
    "name": "Updated Name",
    "email": "updated@example.com",
    "created_at": "2025-07-16 13:45:00",
    "updated_at": "2025-07-16 13:45:00"
  }
}
```

### 5. Delete User
- **Method**: `DELETE`
- **Path**: `/users/:id`
- **Path Parameters**:
  - `id`: User ID (required)
- **Response**:
```json
{
  "success": true
}
```

## Testing

### Using the provided test scripts:

**Windows (PowerShell/CMD):**
```bash
test_users_endpoint.bat
```

**Linux/Mac:**
```bash
chmod +x test_users_endpoint.sh
./test_users_endpoint.sh
```

### Manual testing with curl:

**Get all users:**
```bash
curl -X GET http://localhost:8000/users
```

**Get user by ID:**
```bash
curl -X GET http://localhost:8000/users/1
```

**Create user:**
```bash
curl -X POST http://localhost:8000/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","email":"alice@example.com","password":"secure123"}'
```

**Update user:**
```bash
curl -X PUT http://localhost:8000/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Name","email":"updated@example.com"}'
```

**Delete user:**
```bash
curl -X DELETE http://localhost:8000/users/1
```

## File Structure

The users endpoint follows the GoFrame project structure:

```
api/users/
├── users.go          # Interface definition
└── v1/
    └── users.go      # Request/Response structures

internal/controller/users/
├── users.go              # Controller base (auto-generated)
├── users_new.go          # Controller factory (auto-generated)
└── users_v1_users.go     # Controller implementation
```

## Notes

- Currently uses mock data for demonstration purposes
- In a real application, you would connect to a database
- Validation rules are defined in the request structures
- All endpoints include proper error handling and response formatting
