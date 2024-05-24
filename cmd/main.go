package main

import (
    "price-estimation-api/internal/app"
    "price-estimation-api/internal/config"
)

func main() {
    config.LoadEnv()
    app.Start()
}
