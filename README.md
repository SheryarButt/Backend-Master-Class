# Backend-Master-Class

This repository contains the codes of the Backend Master Class course by TECH SCHOOL.

## [Backend #1] Design DB schema and generate SQL code with dbdiagram.io

- Designed using dbdiagram.io
  - https://dbdiagram.io/d/6272b42f7f945876b6b80516

## [Backend #2] Install & use Docker + Postgres + TablePlus to create DB schema

### Setup Docker

- Get Docker Desktop here:
  - https://www.docker.com/products/docker-desktop/
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
  - https://tableplus.com/windows
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
