package main

import (
    "os"
    "booking-api/controllers"
    "booking-api/worker"
    "github.com/gin-gonic/gin"
)

func main() {
    role := os.Getenv("ROLE")

    if role == "worker" {
        worker.StartWorker()
    } else {
        router := gin.Default()
        router.POST("/api/book", controllers.BookAppointment)
        router.Run(":8080")
    }
}