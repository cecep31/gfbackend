#!/bin/bash

echo "Testing Users Endpoints..."
echo "========================="

echo -e "\n1. GET /users - Get all users"
curl -X GET http://localhost:8000/users

echo -e "\n\n2. GET /users/1 - Get user by ID"
curl -X GET http://localhost:8000/users/1

echo -e "\n\n3. POST /users - Create new user"
curl -X POST http://localhost:8000/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","email":"alice@example.com","password":"secure123"}'

echo -e "\n\n4. PUT /users/1 - Update user"
curl -X PUT http://localhost:8000/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Updated","email":"john.updated@example.com"}'

echo -e "\n\n5. DELETE /users/1 - Delete user"
curl -X DELETE http://localhost:8000/users/1

echo -e "\n\nTesting complete!"
