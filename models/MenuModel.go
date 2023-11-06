package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MenuModel struct {
	ID primitive.ObjectID `bson:"_id"`
	//OrderId   string             `json:"orderId" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Category  string    `json:"category" validate:"required"`
	MenuImage string    `json:"menuImage" validate:"required"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
