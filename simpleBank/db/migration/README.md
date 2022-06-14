## Create migrations 

`migrate.exe create -ext sql -dir db\migration -seq init_schema`

## Upload Migration 

`migrate.exe -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up`