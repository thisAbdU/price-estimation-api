package database

import (
    "database/sql"
    "log"
    "price-estimation-api/internal/config"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func SetupDatabase() {
    var err error
    connStr := config.GetEnv("DATABASE_URL")
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Successfully connected to the database")
}
