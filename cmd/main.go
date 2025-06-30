package main

import (
	"github.com/Arthur-Conti/food_tracker/internal/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RouterRegister(server)
	server.Run(":8080")
}