# Golang Project Assessment by Aqary International Group

## Overview

This document provides an overview of the Golang project, including setup instructions,
database creation steps, and how to run and test the application.

## Setup Instructions

### Creating the Database

First, create a PostgreSQL database named "aqary". Use the following command in the PostgreSQL CLI:

```sql
CREATE DATABASE aqary;
```

### Creating the Users Table

In the `aqary` database, create the `users` table with this SQL command:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL UNIQUE,
    otp VARCHAR(255),
    otp_created_at TIMESTAMP
);
```

## Running the Application

Navigate to the directory containing `main.go` and run the application using:

```bash
go run main.go
```

## API Routes and Testing with Curl

### CreateUser Route

```bash
curl -X POST http://localhost:8080/api/users      -H "Content-Type: application/json"      -d '{"name": "John Doe", "phone_number": "1234567890"}'
```

### GenerateOTP Route

```bash
curl -X POST http://localhost:8080/api/users/generateotp      -H "Content-Type: application/json"      -H "PhoneNumber: 1234567890"      -d '{"phone_number": "1234567890"}'
```

### VerifyOTP Route

```bash
curl -X POST http://localhost:8080/api/users/verifyotp      -H "Content-Type: application/json"      -H "PhoneNumber: 1234567890"      -d '{"phone_number": "1234567890", "otp": "1234"}'
```

## Code Review

Below is a brief overview of each file and its key functionalities:

### `main.go`

Entry point of the application. Initializes the database connection and sets up HTTP routes.

### `otpHandlers.go`

Contains handlers for OTP-related operations.

### `userHandlers.go`

Contains handlers for user-related operations.

### `db_funcs.go`

Includes functions for database interactions.
