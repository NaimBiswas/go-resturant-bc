package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FoodModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name" validate:"required,min=2max=100"`
	Price     float64            `json:"price" validate:"required"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	FoodImage string             `json:"foodImage" validate:"required"`
	MenuId    string             `json:"menuId" validated:"required"`
}
