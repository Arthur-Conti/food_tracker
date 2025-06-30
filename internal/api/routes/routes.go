package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var serviceName = "food-tracker"

func RouterRegister(server *gin.Engine) {

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Server Alive"})
	})

	orderRouter := server.Group(groupName("/order"))
	orderRoutes(orderRouter)
}

func groupName(name string) string {
	return serviceName + name
}