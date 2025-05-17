package models

type ClassInput struct {
 Name string `json:"name" binding:"required"`
 Start string `json:"start_date" binding:"required"`
 End string `json:"end_date" binding:"required"`
 Capacity int `json:"capacity" binding:"required"`
}

type Class struct {
 Name string
 Date string
 Capacity int
}

type BookingInput struct {
 Name string `json:"name" binding:"required"`
 Date string `json:"date" binding:"required"`
}

type Booking struct {
 Name string
 Date string
}