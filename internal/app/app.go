package App

import (
	"log"
	"price-estimation-api/internal/config"
	"price-estimation-api/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct{
    Env *config.Env
    PostgresPool  *pgxpool.Pool 
}
func App() Application {
    app := &Application{}
    app.Env = config.NewEnv()

    // Setup the database connection
    err := database.SetupDatabase(app.Env)
    if err != nil {
        log.Fatalf("Error setting up database: %v", err)
    }
    
    return *app
}

func (app *Application) Close(){
    // Close the database connection
    database.CloseDatabase()
}
