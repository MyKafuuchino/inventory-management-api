package model

type CreateOrderRequest struct {
	Status     string         `json:"status" validate:"omitempty,oneof=pending processed completed canceled"`
	UserID     string         `json:"userId"`
	Quantities map[string]int `json:"quantities" validate:"required"`
	ProductIDs []string       `json:"productIds" validate:"required,min=1,dive"`
}
