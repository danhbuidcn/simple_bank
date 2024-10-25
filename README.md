# Simple Bank Backend Service

## Introduction

This project is a step-by-step guide to designing, developing, and deploying a backend web service from scratch using Golang. Throughout the course, we will build a backend web service for a simple bank, providing APIs to:

- Create and manage bank accounts.
- Record balance changes for each account.
- Perform money transfers between accounts.

## Course Structure
The project is divided into six sections, covering a wide range of backend web development topics:

### 1. Database Design and Development Tools
- Learn how to design a database and generate code for consistent and reliable DB interaction using transactions.
- Understand DB isolation levels and their correct usage in production.
- Use Docker for local development, Git for version control, and GitHub Actions for automated unit testing.

### 2. Building RESTful APIs
- Develop RESTful HTTP APIs using the Gin framework.
- Cover topics such as app configuration, DB mocking for unit tests, error handling, user authentication, and API security using JWT and PASETO tokens.

### 3. Docker and Kubernetes Deployment
- Learn to build a minimal Docker image and deploy it to a Kubernetes cluster on AWS.
- Set up an AWS account, create a production database, manage secrets, and deploy the service using GitHub Actions.
- Secure the service with HTTPS and auto-renew TLS certificates using Let's Encrypt.

### 4. Advanced Backend Topics
- Manage user sessions, build gRPC APIs, and serve both gRPC and HTTP requests.
- Embed Swagger documentation, perform partial record updates, and implement structured logging.

### 5. Asynchronous Processing
- Implement background workers and use Redis as a message queue for asynchronous processing.
- Create and send emails using Gmail's SMTP server, and write unit tests for gRPC services with multiple dependencies.

### 6. Stability and Security
- Improve server stability and security by updating dependencies, securing refresh tokens with cookies, and gracefully shutting down the server.
- This section is continuously updated with new topics.

## Conclusion
This project is designed to be comprehensive, allowing even those with little programming experience to follow along. By the end, you will have gained the confidence and skills to work effectively on your own projects.

## How to run

- Create go project
  ```
  cp .env.sample .env
  docker-compose up --build
  ```

- Migrate db
  ```
  docker exec -it go_app /bin/sh
  # make createdb
  # make migrateup
  # make test
  ```

- Initialize the module and install some packages
  ```
  go mod init simple_bank
  go mod tidy
  ```

##  Install package

```
  docker exec -it go_app /bin/sh
  # go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  # go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
  # go get -u github.com/gin-gonic/gin
  # go get github.com/spf13/viper
  # go install github.com/golang/mock/mockgen@v1.6.0
  # go get github.com/golang/mock/gomock
```

### golang-migrate

- [docs](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### sqlc

- [docs](https://docs.sqlc.dev/en/stable/overview/install.html) 
- You can also use [gorm](https://gorm.io/docs/) instead.

```
  docker exec -it go_app /bin/sh

  sqlc version
  sqlc generate 
```

### testify

[testify](https://github.com/stretchr/testify)

## Commit History

### Section 1: Working with database [Postgres + SQLC]
  - **feat: initialize project and connect db**: 
    - Initialized the project structure.
    - Connected to the database using Docker for containerization.

  - **feat: setup sqlc and implement CRUD with sqlc**: 
    - Set up sqlc for generating Go code from SQL queries.
    - Implemented CRUD operations using sqlc.

  - **feat: add comprehensive tests for account operations**: 
    - Wrote extensive tests for various account operations to ensure reliability.

  - **feat: add comprehensive tests for account operations**: 
    - Added more comprehensive tests for account operations.

  - **feat: add CRUD entry and transfer**: 
    - Added Create, Read, Update, Delete (CRUD) operations for entries and transfers.

  - **feat: implement database transaction, transfers transaction**: 
    - Implemented database transactions.
    - Implemented transfer transactions.

  - **feat: eliminate duplicate code and handle deadlocks**: 
    - Removed duplicate code.
    - Added handling for database deadlocks.

  - **docs: deeply understand transaction isolation levels & read phenomena**: 
    - Added documentation to deeply understand transaction isolation levels.
    - Explained read phenomena.

  - **feat: setup github actions for project**: 
    - Set up GitHub Actions for continuous integration and deployment.

### Section 2: Building RESTful HTTP JSON API [Gin + JWT + PASETO]
  - **feat: implement RESTful HTTP API in go using Gin**
    - Developed RESTful HTTP APIs using the Gin framework.
    - Implemented endpoints for creating accounts.
  - **feat: load config from file & environment variables with viper**
    - [viper](https://github.com/spf13/viper)
    - Load config from `util/config.go` by viper
  - **feat: mock DB for testing HTTP API in Go and achieve 100% coverage** (use [gomock](https://github.com/golang/mock))
    - Isolate tests data to avoid conflicts
    - Reduce a lot of time talking to the database
    - Easy setup edge cases: unexpected errors
    - Command: `make mock`
  - **feat: implement transfer money API with a custom params validator**
    - Added logic to validate currency and balance before performing transfer

### Section 3: Deploying the application to production [Docker + Kubernetes + AWS]

### Section 4: Advance backend topics [Sessions + gRPC]

### Section 5: Asynchronous processing with background workers [Asynq + Redis]

### Section 6: Improve the stability and security of the server [PGX + RBAC + CORS]
