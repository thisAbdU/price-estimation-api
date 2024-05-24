package App

import (
    "price-estimation-api/internal/config"
    "price-estimation-api/internal/database"

)

func App() {
    // Load environment variables
    env := config.NewEnv()

    // Set up the database
    if err := database.SetupDatabase(env); err != nil {
        panic(err)
    }

}

func Close() {
    // Close the database connection
    database.CloseDatabase()
}
