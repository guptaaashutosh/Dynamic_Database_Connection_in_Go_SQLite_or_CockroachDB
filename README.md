# Dynamic Database Connection with Go

This repository demonstrates how to create a dynamic database connection in a Go application by changing a single flag in the environment configuration. It includes implementations for both SQLite and CockroachDB.

## Overview

In this project, we've implemented a dynamic database connection mechanism that allows developers to easily switch between SQLite and CockroachDB with minimal changes. This is achieved by utilizing environment variables to specify the desired database type.

### Features

- Dynamic database connection switching
- SQLite and CockroachDB implementations
- Clear separation of concerns for database interactions

## Requirements

To run this project, you need the following:

- Go programming language
- `go-sqlite3` package for SQLite support (optional)
- CockroachDB (if using CockroachDB)

## Getting Started

Follow these steps to set up and run the project:

1. Clone the repository:

   ```bash
   git clone https://github.com/guptaaashutosh/Dynamic_Database_Connection_in_Go_SQLite_or_CockroachDB.git
   cd Dynamic_Database_Connection_in_Go_SQLite_or_CockroachDB
   ```


2. Install Dependencies:
    ```bash
    go mod tidy
    ```
    or,

3. If using SQLite, install the go-sqlite3 package:
    ```bash
    go get github.com/mattn/go-sqlite3
    ```
4. Install the CockroachDB driver for Go (refer to official documentation for specific instructions).

    https://www.cockroachlabs.com/docs/v24.1/install-cockroachdb-windows.html

5. Create `.config` folder :   create `.env` file inside it and configure as per `.env.example` file.
```
Dynamic_Database_Connection_in_Go_SQLite_or_CockroachDB/
├── .config/
│   └── .env
├── .env.example
```

6. Set the `DB_CONNECTION` type  environment variable to either "sqlite" or "cockroachdb" as per your requirement.
<br><br>
The dynamic database connection is controlled by the `DB_CONNECTION` environment variable. Set this variable to `sqlite` or `cockroachdb` to choose the database type.
    ```bash
    DB_CONNECTION="sqlite"
    ```

7. Run the Application:
```bash
go run cmd/main.go
```

### Benefits:

- **Flexibility:** Switch between databases based on your needs (development, testing, production).
- **Reusability:** Database logic remains consistent regardless of the underlying database.
- **Maintainability:** Code is easier to understand and maintain with a clear separation of concerns.

### Further Considerations:
- This example provides a basic structure. Adapt it to your specific database interactions and error handling.
- Consider adding support for additional database types if needed.


<br><br>

This readme provides a starting point for creating a dynamic and flexible database connection system in your Go applications.






