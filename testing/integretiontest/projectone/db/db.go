package db

import (
"database/sql"
 	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error){

connStr:="postgres://postgres:5511@localhost:5432/testdb?sslmode=disable"

return sql.Open("postgres", connStr)

}