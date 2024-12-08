# Go Registration API with Gin and MongoDB

A simple user registration API built with Go using the Gin framework and MongoDB. The API securely hashes user passwords with bcrypt before storing them in the database.

## Prerequisites

- **Go**: [Install Go](https://golang.org/dl/)
- **MongoDB**: [Install MongoDB](https://www.mongodb.com/docs/manual/installation/) or use MongoDB Atlas.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/go-registration-api.git
   cd go-registration-api
   ```

2. Install dependencies:

   ```bash
   go get github.com/gin-gonic/gin
   go get go.mongodb.org/mongo-driver/mongo
   go get golang.org/x/crypto/bcrypt
   ```

## Running the Application

1. Ensure MongoDB is running locally or use a remote instance.
2. Start the application:

   ```bash
   go run main.go
   ```

3. The server will be available at `http://localhost:8080`.

## API Endpoints

- **POST /register**: Register a new user.
  
  Request body (JSON):
  ```json
  {
    "name": "John Doe",
    "email": "johndoe@example.com",
    "password": "password123"
  }
  ```

  Response (Success):
  ```json
  {
    "status": true,
    "message": "User successfully registered",
    "data": {
      "name": "John Doe",
      "email": "johndoe@example.com"
    }
  }
  ```

  Response (Error):
  ```json
  {
    "status": false,
    "message": "Failed to insert data"
  }
  ```