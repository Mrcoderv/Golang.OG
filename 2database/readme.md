# Database Access in Go

A database stores and manages application data. In Go, an ORM (Object-Relational
Mapper) helps map Go structs to database tables and reduces the amount of SQL
code that must be written by hand.

This example uses [GORM](https://gorm.io/). Bun is another Go ORM, but it is a
separate library and is not used in this project.

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


