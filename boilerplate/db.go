package boilerplate

const DbConnPostgresGoFileData = `
package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetDBConn() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
}
`
const SQLCYAMLPostgresFileData = `
version: "2"

sql:
  - engine: "postgresql"
    schema: "sql/schema"
    queries: "sql/queries"
    gen:
      go:
        package: "db"
        out: "internal/db"
        sql_package: "pgx/v5"
        emit_pointers_for_null_types: true
`

const DbConnSQLiteGoFileData = `
package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetDBConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
`

const SQLCYAMLSQLiteFileData = `
version: "2"

sql:
  - engine: "sqlite"
    schema: "sql/schema"
    queries: "sql/queries"
    gen:
      go:
        package: "db"
        out: "internal/db"
        sql_package: "database/sql"
`
