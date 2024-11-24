
---

# Expense Tracker API

## Overview

Expense Tracker API is a backend service built with **Go** and the **Gin** framework. It is designed to help users manage their transactions securely. The application uses **MongoDB** as the database and includes endpoints for user authentication and transaction management.

---

## Features

- **User Authentication**:
  - JWT-based authentication for secure access to protected routes.
  
- **Transaction Management**:
  - Create, read, update, and delete transactions.

- **RESTful API**:
  - Well-structured public and protected endpoints.

- **MongoDB Integration**:
  - Fully integrated with MongoDB for transaction and user data.

---

## Requirements

Before running the project, ensure the following are installed:

- **Go**: `v1.22` or later
- **MongoDB**: A running MongoDB instance
- **Git**: To clone the repository

---

## Setup Instructions

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/expense-tracker-go.git
   cd expense-tracker-go
   ```

2. **Install Dependencies**:
   Install Go modules using:
   ```bash
   go mod tidy
   ```

3. **Set Environment Variables**:
   Create a `.env` file in the root directory with the following variables:
   ```
   MONGO_URI=mongodb://localhost:27017
   DB_NAME=ExpenseTracker
   JWT_SECRET=your_secret_key
   PORT=8080
   ```

4. **Run the Application**:
   Start the server using:
   ```bash
   go run main.go
   ```

5. **Access the API**:
   - Public endpoints: `http://localhost:8080/api/v1/public`
   - Protected endpoints: `http://localhost:8080/api/v1/protected` (require JWT tokens)

---

## Project Structure

```
Expense-Tracker-go/
├── controllers/         # Handles HTTP requests and responses
├── middleware/          # Middleware for authentication (e.g., JWT validation)
├── models/              # Data models (e.g., User, Transaction)
├── repository/          # Handles database queries and interactions
├── routes/              # API route definitions
├── services/            # Business logic and data validation
├── utils/               # Utility functions (e.g., database connection)
├── .env                 # Environment variables (not included in version control)
├── go.mod               # Dependency definitions
├── go.sum               # Dependency checksums
├── main.go              # Entry point of the application
└── README.md            # Documentation
```

---

## API Endpoints

### Public Endpoints

| Method | Endpoint         | Description           |
|--------|------------------|-----------------------|
| POST   | `/api/v1/public/signup` | Register a new user   |
| POST   | `/api/v1/public/login`  | Login a user          |

### Protected Endpoints

| Method | Endpoint                   | Description                   |
|--------|----------------------------|-------------------------------|
| GET    | `/api/v1/protected/`       | Get all user transactions     |
| GET    | `/api/v1/protected/latest` | Get the latest 4 transactions |
| GET    | `/api/v1/protected/:id`    | Get transaction by ID         |
| POST   | `/api/v1/protected/`       | Add a new transaction         |
| PUT    | `/api/v1/protected/:id`    | Update a transaction by ID    |
| DELETE | `/api/v1/protected/:id`    | Delete a transaction by ID    |

---

## Dependencies

Here are the main libraries used in this project:

| Library                                     | Version   | Description                                                                                      |
|---------------------------------------------|-----------|--------------------------------------------------------------------------------------------------|
| `github.com/gin-gonic/gin`                  | v1.10.0   | High-performance web framework for building REST APIs.                                           |
| `go.mongodb.org/mongo-driver`               | v1.17.1   | MongoDB driver for Go, used for database operations.                                             |
| `github.com/golang-jwt/jwt/v5`              | v5.2.1    | JWT library for creating and validating authentication tokens.                                   |
| `github.com/go-playground/validator/v10`    | v10.23.0  | Validation library for validating user inputs and struct fields.                                 |
| `github.com/sirupsen/logrus`                | v1.9.3    | Structured logger for debugging and monitoring.                                                 |
| `github.com/joho/godotenv`                  | v1.5.1    | Loads environment variables from a `.env` file.                                                 |
| `github.com/goccy/go-json`                  | v0.10.3   | High-performance JSON encoder/decoder for faster responses.                                      |

---

## Running Tests

To run unit tests (if implemented):
```bash
go test ./...
```

---

## Logging

The server logs important events such as database connectivity, server startup, and errors to the console. Structured logging is implemented using the **Logrus** library.

---

## Contributions

Contributions are welcome! Follow these steps to contribute:

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add feature name"
   ```
4. Push to your branch:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request.

---

## License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

---

## Contact

For questions or support, please contact:

- **Email**: kaweesha.mr@gmail.com
- **GitHub**: [kaweesha-mr](https://github.com/kaweesha-mr)

---
