package dto

import "github.com/google/uuid"

type Order struct {
	ID          uuid.UUID
	UserName    string
	ProductName string
	Status      string
}
