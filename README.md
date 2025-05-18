# Glofox API (Assignment)

This is a simple in-memory REST API built using the [Gin](https://github.com/gin-gonic/gin) framework in Go. It simulates a studio/gym class booking system for the Glofox assignment.

## Features

- Create fitness classes (`/classes`)
- Book a spot in a class (`/bookings`)
- In-memory storage (no external DB)
- Basic validations

---

## Prerequisites

- Go 1.18+
- curl or Postman for API testing

---

## Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/vicioas/Glofox.git
cd Glofox
go mod tidy
go run main.go
