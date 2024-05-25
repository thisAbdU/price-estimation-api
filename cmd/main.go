package main

import (
	"fmt"
	"price-estimation-api/internal/app"
	"time"
)

func main() {
    application := App.App()
    defer application.Close()

    env := application.Env

    timeout := time.Duration(env.ContextTimeout) * time.Second
    fmt.Println("Server running on port 8080")
    
   

}
