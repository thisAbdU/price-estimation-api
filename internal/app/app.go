package main

import (
    "github.com/gin-gonic/gin"
    "price-estimation-api/internal/handlers"
    "price-estimation-api/internal/config"
    "price-estimation-api/internal/database"
)

func main() {
    // Load environment variables
    config.LoadEnv()
    
    // Set up the database
    database.SetupDatabase()
    
    router := gin.Default()
    
    router.POST("/estimates", handlers.CreateEstimate)
    router.GET("/estimates", handlers.GetEstimates)
    router.GET("/estimates/:id", handlers.GetEstimate)
    router.PATCH("/estimates/:id", handlers.UpdateEstimate)
    router.DELETE("/estimates/:id", handlers.DeleteEstimate)

    // Start the server
    port := config.GetEnv("PORT")
    if port == "" {
        port = "8080"
    }
    router.Run(":" + port)
}
