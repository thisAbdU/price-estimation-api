package database

import (
	"context"
	"database/sql"
	"log"
	"price-estimation-api/internal/config"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func SetupDatabase(env *config.Env) error {
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()

   dbUrl := "host=" + env.DB_HOST + " port=" + env.DB_PORT + " user=" + env.DB_USER + " dbname=" + env.DB_NAME + " password=" + env.DB_PASSWORD + " sslmode=disable"
    db, err := sql.Open("postgres", dbUrl)
    if err != nil {
        return err
    }

    err = db.PingContext(ctx)
    if err != nil {
        return err
    }

    DB = db
    log.Println("Database connection established")
    return nil

}

func CloseDatabase() {
    if DB != nil {
        DB.Close()
        log.Println("Database connection closed")
    }
}