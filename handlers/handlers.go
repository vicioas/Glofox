package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glofox/models"
	"github.com/glofox/store"
	"github.com/glofox/utils"
)

func CreateClass(c *gin.Context) {
 var input models.ClassInput
 if err := c.ShouldBindJSON(&input); err != nil {
 c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input" + err.Error()})
 return
 }

 dates, err := utils.GenerateDateRange(input.Start, input.End)
 if err != nil {
 c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date range"})
 return
 }

 for _, date := range dates {
 store.Classes[date] = models.Class{
 Name: input.Name,
 Date: date,
 Capacity: input.Capacity,
 }
 }

 c.JSON(http.StatusCreated, gin.H{"message": "Class created"})
}

func CreateBooking(c *gin.Context) {
 var input models.BookingInput
 if err := c.ShouldBindJSON(&input); err != nil {
 c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
 return
 }

 if _, ok := store.Classes[input.Date]; !ok {
 c.JSON(http.StatusNotFound, gin.H{"error": "No class on this date"})
 return
 }

 booking := models.Booking{Name: input.Name, Date: input.Date}
 store.Bookings[input.Date] = append(store.Bookings[input.Date], booking)

 c.JSON(http.StatusCreated, gin.H{"message": "Booking successful"})
 
}
