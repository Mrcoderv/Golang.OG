# Database Access in Go

A database stores and manages application data. In Go, an ORM (Object-Relational
Mapper) helps map Go structs to database tables and reduces the amount of SQL
code that must be written by hand.

This example uses [GORM](https://gorm.io/). Bun is another Go ORM, but it is a
separate library and is not used in this project.

## PostgreSQL connection

The CRUD example reads its connection string from `DATABASE_URL`. Set it
before running the program, using the PostgreSQL role and password configured
on your machine:

```bash
export DATABASE_URL='postgres://mrrv:YOUR_PASSWORD@localhost:5432/go_crud?sslmode=disable'
go run .
```

Do not commit the connection string or password to source control. If the
database role does not exist, create it in PostgreSQL or use an existing role.

## Connection flow

```text
Go application
        |
GORM database driver (for example, gorm.io/driver/postgres)
        |
PostgreSQL database
```

Install the PostgreSQL driver and GORM:

```bash
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

Open a GORM connection with `gorm.Open`:

```go
db, err := gorm.Open(
      postgres.Open("postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"),
      &gorm.Config{},
)
if err != nil {
      log.Fatal(err)
}
```

GORM opens and configures the database connection. To explicitly check that
the underlying database is reachable, get the `database/sql` connection and
call `Ping`:

```go
sqlDB, err := db.DB()
if err != nil {
      log.Fatal(err)
}
if err := sqlDB.Ping(); err != nil {
      log.Fatal(err)
}
```

GORM can then perform common operations such as creating records, querying
rows, updating records, deleting records, and migrating tables.

