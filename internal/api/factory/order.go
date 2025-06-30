package factory

import orderservice "github.com/Arthur-Conti/food_tracker/internal/services/order"

func OrderServiceFactory() *orderservice.OrderService{

	return orderservice.NewOrderService()
}