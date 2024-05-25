package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"price-estimation-api/internal/app"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
    application := App.App()
    defer application.Close()

    env := application.Env
    router := gin.Default()

   
    server := &http.Server{
        Addr:         fmt.Sprintf(":%s", env.DB_PORT),
        Handler:      router,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    go func() {
        if err := server.ListenAndServe(); err != nil {
            log.Fatalf("Server failed to start: %v", err)
        }
    }()

    
    fmt.Println("Server running on port 8080")
    
    // Graceful shutdown
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
    
}
