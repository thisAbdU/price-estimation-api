package database

import (
    "database/sql"
    "log"
    "price-estimation-api/internal/config"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func SetupDatabase() error {
    connStr := config.GetEnv("DATABASE_URL")
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return err
    }

    if err := db.Ping(); err != nil {
        db.Close()
        return err
    }

    DB = db
    log.Println("Successfully connected to the database")
    return nil
}

func CloseDatabase() {
    if DB != nil {
        DB.Close()
        log.Println("Database connection closed")
    }
}
