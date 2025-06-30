package ordercontroller

import (
	"net/http"

	"github.com/Arthur-Conti/food_tracker/internal/api/controllers"
	"github.com/Arthur-Conti/food_tracker/internal/domain/dto"
	orderservice "github.com/Arthur-Conti/food_tracker/internal/services/order"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	svc *orderservice.OrderService
}

func NewOrderController(svc *orderservice.OrderService) *OrderController {
	return &OrderController{
		svc: svc,
	}
}

func (oc *OrderController) CreateHandler(c *gin.Context) {
	var input CreateOrderModel

	if err := c.ShouldBindJSON(&input); err != nil {
		controllers.FailResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := oc.svc.Create(&dto.Order{
		UserName:    input.UserName,
		ProductName: input.ProductName,
	})
	if err != nil {
		controllers.FailResponse(c, http.StatusBadRequest, err.Error())
	}

	response := CreateSuccessfully{
		ID:      id,
		Message: "Order successfully created",
	}

	controllers.SuccessResponse(c, http.StatusOK, response)
}
