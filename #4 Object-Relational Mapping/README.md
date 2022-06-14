# [Backend #4] Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc

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

## Usage

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
