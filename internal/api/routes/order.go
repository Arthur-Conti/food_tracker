package routes

import (
	ordercontroller "github.com/Arthur-Conti/food_tracker/internal/api/controllers/order"
	"github.com/Arthur-Conti/food_tracker/internal/api/factory"
	"github.com/gin-gonic/gin"
)

func orderRoutes(group *gin.RouterGroup) {
	service := factory.OrderServiceFactory()
	controller := ordercontroller.NewOrderController(service)

	group.POST("/", controller.CreateHandler)
}