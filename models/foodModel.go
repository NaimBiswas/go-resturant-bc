package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FoodModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `json:"name" bson:"name" validate:"required,min=2,max=100"`
	Price     float64            `json:"price" validate:"required" bson:"price"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	FoodImage string             `json:"foodImage" validate:"required" bson:"foodImage"`
	MenuId    string             `json:"menuId" validated:"required" bson:"menuId"`
	IsActive  bool               `json:"isActive" bson:"isActive" bson:"isActive"`
}
