# Backend Golang Project

## Project Overview
This project is a backend application built with Golang that provides an API for retrieving cat facts. It utilizes Gorm as the Object Relational Mapping (ORM) library and MySQL as the database. The project includes Docker configuration for easy deployment.

## Built With
- [Golang](https://golang.org/): The programming language used for backend development.
- [Gorm](https://gorm.io/): A powerful Object Relational Mapping (ORM) library for Golang.
- [Goose](https://github.com/pressly/goose): A database migration tool for Golang applications.
- [MySQL](https://www.mysql.com/): The relational database used for storing data.
- [Docker](https://www.docker.com/): A containerization platform used for packaging the application and its dependencies.

## Project Directory
```
- main.go
- api/
  - handlers/
    - api_handlers.go
  - routes/
    - api_routes.go
- database/
  - database.go
    - migrations/
      - 20230627174850_init_table
- models/
  - fact.go
- utils/
  - errors.go
```

## Getting Started
To run this project locally, follow these steps:

1. Make sure the database is set up and running on your local machine.
2. Run the database migrations using the Goose tool:
```
goose up
```
3. Run the project using Go:
```
go run main.go
```

or
4. Build and run the project using Docker:
```
docker build -t go-fact .
docker run -p 8000:8000 go-fact
```

## API Documentation and Integrations
The API exposes the following endpoint:
- GET `/api/facts`

This endpoint retrieves cat facts from the free open-source API [Cat Fact Ninja](https://catfact.ninja/).

## Configuration
The project can be configured by modifying the following constants in the code:

```go
const (
 DatabaseHost     = "localhost"
 DatabaseUsername = "root"
 DatabasePassword = "toor"
 DatabasePort     = "3306"
 DatabaseName     = "cats"
)
```
Make sure to update the database configuration according to your local setup.