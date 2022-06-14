# [Backend #3] How to write & run database migration in Golang

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

## Create migrations

- Initialize migration using:
  - `migrate.exe create -ext sql -dir db\migration -seq init_schema`
- Move the generated SQL code from dbdiagram.io to:
  - \db\migration\000001_init_schema.up.sql
- Add a rollback plan in case of failure to:
  - \db\migration\000001_init_schema.down.sql

## Upload Migration

- Upload migration using:
  - `migrate.exe -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up`
  - Or use `make migrateup`

## Rollback Migration

- Rollback migration using:
  - `migrate.exe -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down`
  - Or use `make migratedown`
