package boilerplate

const DbConnGoFileData = `
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
const SQLCYAMLFileData = `
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
