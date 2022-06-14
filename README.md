# Backend-Master-Class

This repository contains the codes of the Backend Master Class course by TECH SCHOOL.

## [Backend #1] Design DB schema and generate SQL code with dbdiagram.io

- Designed using dbdiagram.io
  - <https://dbdiagram.io/d/6272b42f7f945876b6b80516>

## [Backend #2] Install & use Docker + Postgres + TablePlus to create DB schema

### Setup Docker

- Get Docker Desktop here:
  - <https://www.docker.com/products/docker-desktop/>
  - Make sure you have WSL enabled if on windows.

### Setup Postgres

- Setup Postgres using:
  - docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
  - Or use `make postgres`
- Create DB using:
  - docker exec -it postgres12 createdb --username=root --owner=root simple_bank
  - Or use `make createdb`
- Drop DB using:
  - docker exec -it postgres12 dropdb simple_bank
  - Or use `make dropdb`

### Setup TablePlus

- Get TablePlus here:
  - <https://tableplus.com/windows>
- Setup TablePlus using:
  - Choose "Create a new connection"
  - Enter the following:
    - Host: localhost
    - Port: 5432
    - Database: simple_bank
    - Username: root
    - Password: secret
    - Click "Connect"
  
## [Backend #3] How to write & run database migration in Golang

- We'll be using golang migrate for database migrations that can be found here:
  - `https://github.com/golang-migrate/migrate`
- To install, head over to CLI Documentation:
  - `https://github.com/golang-migrate/migrate/tree/master/cmd/migrate`
- For windows, use scoop to install golang-migrate:
  - Install scoop using
    - `https://github.com/ScoopInstaller/Install#readme`
  - Install golang-migrate using:
    - `scoop install golang-migrate`
- Verify installation using
  - `migrate.exe --version`

### Create migrations

- Initialize migration using:
  - `migrate.exe create -ext sql -dir db\migration -seq init_schema`
- Move the generated SQL code from dbdiagram.io to:
  - \db\migration\000001_init_schema.up.sql
- Add a rollback plan in case of failure to:
  - \db\migration\000001_init_schema.down.sql

### Upload Migration

- Upload migration using:
  - `migrate.exe -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up`
  - Or use `make migrateup`

### Rollback Migration

- Rollback migration using:
  - `migrate.exe -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down`
  - Or use `make migratedown`

## [Backend #4] Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc

We have 4 options for Object-Relational Mapping (ORM) in golang.

- Database/sql (Go standard library)
- GORM (Go ORM) : `gorm.io`
- SQLX : `github.com/jmoiron/sqlx`
- SQLC : `sqlc.dev`

We'll move forward with SQLC.

- Use go get to install sqlc.
  - `go get github.com/kyleconroy/sqlc/cmd/sqlc`
- Or use the pre-built libraries from:
  - `https://docs.sqlc.dev/en/latest/overview/install.html#go-install`
  - Add binary to your PATH.
  - Verify that the binary is working using `sqlc version`.

### Usage

- Add sqlc.yaml to your project root with the following content:

```yaml
  version: "1"
project:
    id: "1"
packages:
  - name: "db"
    path: "./db/sqlc"
    queries: "./db/query/"
    schema: "./db/migration/"
    engine: "postgresql"
    emit_json_tags: true
    emit_prepared_queries: false
    emit_interface: false
    emit_exact_table_names: false 
```

- Run the following to initialize a new project.
  - `sqlc init`
  - or if you are on windows, use from the project root:
    - `docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc init`
- Go to <https://docs.sqlc.dev/en/stable/> and use HOW-TO guides to write your queries and place them in db/query/ folder.
- Generate Go code using:
  - `sqlc generate`
  - or if you are on windows, use from the project root:
    - `docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc generate`

## [Backend #5] Write Golang unit tests for database CRUD with random data

For unit tests, we'll be using Go's testing framework with **Testify**.

- Get Testify here:
  - `go get github.com/stretchr/testify'
- We'll be using *testify/require* to verify our tests.
- The `require` package provides same global functions as the assert package, but instead of returning a boolean result they terminate current test.
- Execute unit-tests using:
  - `go test -v -cover ./...`
  - or use `make test`
  