package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/glofox/handlers"
)

func main(){
	fmt.Println("Test")
	server := gin.Default()
	server.POST("/classes",handlers.CreateClass)

	server.Run(":8080")
}