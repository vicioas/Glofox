package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/glofox/models"
	"github.com/glofox/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Setup router for tests
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/classes", CreateClass)
	r.POST("/bookings", CreateBooking)
	return r
}

func TestCreateClass(t *testing.T) {
	router := setupRouter()

	input := models.ClassInput{
		Name:     "Yoga",
		Start:    "2025-06-01",
		End:      "2025-06-03",
		Capacity: 15,
	}
	body, _ := json.Marshal(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/classes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), "Class created")
}

func TestCreateBooking_Success(t *testing.T) {
	router := setupRouter()

	// First create the class
	store.Classes = map[string]models.Class{} // clear before test
	store.Bookings = map[string][]models.Booking{} // clear bookings

	_ = httptest.NewRecorder()
	classInput := models.ClassInput{
		Name:     "HIIT",
		Start:    "2025-07-01",
		End:      "2025-07-01",
		Capacity: 5,
	}
	classBody, _ := json.Marshal(classInput)
	classReq, _ := http.NewRequest("POST", "/classes", bytes.NewBuffer(classBody))
	classReq.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(httptest.NewRecorder(), classReq)

	// Book a class
	booking := models.BookingInput{
		Name: "Bob",
		Date: "2025-07-01",
	}
	bookBody, _ := json.Marshal(booking)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/bookings", bytes.NewBuffer(bookBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), "Booking successful")
}

func TestCreateBooking_NoClass(t *testing.T) {
	router := setupRouter()
	store.Classes = map[string]models.Class{} // clear before test

	booking := models.BookingInput{
		Name: "John",
		Date: "2025-08-01",
	}
	bookBody, _ := json.Marshal(booking)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/bookings", bytes.NewBuffer(bookBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Contains(t, w.Body.String(), "No class on this date")
}



