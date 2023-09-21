# Go Backend Banking w/ gRPC, Gin, Postgres, Docker, K8S, AWS, GHA


## DB


### CMDS

- Running Postgres container `docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres:12-alpine`
- Exec Docker container `docker exec -it postgres12 psql -U root`
- Postgres Shell `docker exec -it postgres12 /bin/sh`
- Create db `createdb --username=root --owner=root simple_bank`
- Access PG db w/ Docker `docker exec -it postgres12 psql -U root simple_bank`

### DB Tools
- [dbdiagram.io](http://dbdiagram.io) for SQL generation model
- [Tableplus](http://tableplus.com)  DB management
- [sqlc](http://sqlc.dev) SQL to CRUD Go code generator. Creates files in db/sqlc of data structs

### Go Packages
- [pq](https://github.com/lib/pq) Postgres Driver 
- [testify](https://github.com/stretchr/testify) Test Assertions 
- [golang-migrate](https://github.com/golang-migrate/migrate) library DB migrations

Generate migration: `migrate create -ext sql -dir db/migration -seq init_schema`
Run migration: `migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up`

