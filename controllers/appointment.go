package controllers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "booking-api/kafka"
    "booking-api/models"
)

func BookAppointment(c *gin.Context) {
    var appointment models.Appointment
    if err := c.ShouldBindJSON(&appointment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    appointment.ID = primitive.NewObjectID()
    appointment.CreatedAt = time.Now()

    message, err := models.ToJSON(appointment)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize appointment"})
        return
    }

    kafka.Produce("bookings", message)

    c.JSON(http.StatusOK, gin.H{"message": "Booking request sent successfully"})
}