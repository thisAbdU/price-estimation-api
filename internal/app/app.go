package main

import (
    "github.com/gin-gonic/gin"
    "price-estimation-api/internal/handlers"
    "price-estimation-api/internal/config"
    "price-estimation-api/internal/database"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // Load environment variables
    config.LoadEnv()

    // Set up the database
    if err := database.SetupDatabase(); err != nil {
        // Handle database setup error
        panic(err)
    }

    // Initialize handlers
    estimateRepo := database.DB
    estimateHandler := handlers.NewHandlers(estimateRepo)

    // Initialize Gin router
    router := gin.Default()

    // Define routes
    router.POST("/estimates", estimateHandler.CreateEstimate)
    router.GET("/estimates", estimateHandler.GetEstimates)
    router.GET("/estimates/:id", estimateHandler.GetEstimate)
    router.PATCH("/estimates/:id", estimateHandler.UpdateEstimate)
    router.DELETE("/estimates/:id", estimateHandler.DeleteEstimate)

    // Start the server
    port := config.GetEnv("PORT")
    if port == "" {
        port = "8080"
    }
    go func() {
        if err := router.Run(":" + port); err != nil {
            panic(err)
        }
    }()

    // Handle graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    if err := database.CloseDatabase(); err != nil {
        // Handle database close error
        panic(err)
    }
}
