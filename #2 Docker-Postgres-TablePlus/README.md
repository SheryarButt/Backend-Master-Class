# [Backend #2] Install & use Docker + Postgres + TablePlus to create DB schema

## Setup Docker

- Get Docker Desktop here:
  - https://www.docker.com/products/docker-desktop/
  - Make sure you have WSL enabled if on windows.

## Setup Postgres

- Setup Postgres using:
  - docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
  - Or use `make postgres`
- Create DB using:
  - docker exec -it postgres12 createdb --username=root --owner=root simple_bank
  - Or use `make createdb`
- Drop DB using:
  - docker exec -it postgres12 dropdb simple_bank
  - Or use `make dropdb`

## Setup TablePlus

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
  