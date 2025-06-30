package ordercontroller

import "github.com/google/uuid"

type CreateOrderModel struct {
	UserName    string `json:"user_name"`
	ProductName string `json:"product_name"`
}

type CreateSuccessfully struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}
