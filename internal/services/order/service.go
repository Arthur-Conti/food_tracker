package orderservice

import (
	"github.com/Arthur-Conti/food_tracker/internal/domain/dto"
	"github.com/google/uuid"
)

type OrderService struct {

}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (os *OrderService) Create(input *dto.Order) (uuid.UUID, error) {
	
	return uuid.New(), nil
}