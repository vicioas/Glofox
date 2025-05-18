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
- Alternatively you can install the REST Client extension on VSCode and test the api using the files in the api-test folder

---

## API Endpoints

### /classes

Request

```bash

{
 "name": "Pilates",
 "start": "2025-06-01",
 "end": "2025-06-10",
 "capacity": 10
}

```

Response

```bash

{
 "message": "Class created"
}

```

### /bookings

Request 

```bash

{
 "name": "Alice",
 "date": "2025-06-01"
}

```

Response

```bash

{
 "message": "Booking successful"
}

```

## Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/vicioas/Glofox.git
cd Glofox
go mod tidy
go run main.go

```

### 2. Test the endpoints

You can just click on the send request button on the files in the api-test folder if you have the REST Client extension installed or you can run the following curl commands.

```bash
curl -X POST http://localhost:8080/classes \
 -H "Content-Type: application/json" \
 -d '{
 "name": "Pilates",
 "start": "2025-06-01",
 "end": "2025-06-05",
 "capacity": 10
}'


curl -X POST http://localhost:8080/bookings \
 -H "Content-Type: application/json" \
 -d '{
 "name": "Alice",
 "date": "2025-06-02"
}'

```

