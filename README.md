# Description

This is golang clean architecture.

## Architecture

![Clean Architecture](architecture.png)

1. External system perform request (HTTP, gRPC, Messaging, etc)
2. The Delivery creates various Model from request data
3. The Delivery calls Use Case, and execute it using Model data
4. The Use Case create Entity data for the business logic
5. The Use Case calls Repository, and execute it using Entity data
6. The Repository use Entity data to perform database operation
7. The Repository perform database operation to the database
8. The Use Case create various Model for Gateway or from Entity data
9. The Use Case calls Gateway, and execute it using Model data
10. The Gateway using Model data to construct request to external system 
11. The Gateway perform request to external system (HTTP, gRPC, Messaging, etc)

## Tech Stack

- Golang : https://github.com/golang/go
- MySQL (Database) : https://github.com/mysql/mysql-server

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus
- Confluent Kafka Golang : https://github.com/segmentio/kafka-go

## Configuration

All configuration is in `config.json` file.

## API Spec

All API Spec or postman collection is in `api` folder.

## Database Migration

All database migration is in `db/migrations` folder.

### Install Migrate PostgreSQL

```shell
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Create Migration

```shell
migrate create -ext sql -dir db/migrations create_table_xxx
```

### Run Migration

```shell
migrate -database "postgres://postgres:mysecretpassword@localhost:5432/assessment_go_source_code_muhammad_aditya?sslmode=disable" -path db/migrations up
```

## Run Application

### Run unit test

```bash
go test -v ./test/
```

### Run web server

```bash
go run cmd/web/main.go
```
