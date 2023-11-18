package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MenuModel struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	//OrderId   string             `json:"orderId" validate:"required"`
	Name      string    `json:"name" validate:"required" bson:"name"`
	Category  string    `json:"category" validate:"required" bson:"category"`
	MenuImage string    `json:"menuImage" validate:"required" bson:"menuImage"`
	StartDate string    `json:"startDate" bson:"startDate"`
	EndDate   string    `json:"endDate" bson:"endDate"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	IsActive  bool      `bson:"isActive" json:"isActive" validate:"required"`
	CreateBy  string    `bson:"createBy" json:"createBy"`
	UpdatedBy string    `bson:"updateBy" json:"updateBy"`
}

type UpdateMenuModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `json:"name" validate:"required" bson:"name"`
	Category  string             `json:"category" validate:"required" bson:"category"`
	MenuImage string             `json:"menuImage" validate:"required" bson:"menuImage"`
	StartDate string             `json:"startDate" bson:"startDate"`
	EndDate   string             `json:"endDate" bson:"endDate"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	IsActive  bool               `bson:"isActive" json:"isActive" validate:"required"`
	UpdatedBy string             `bson:"updateBy" json:"updateBy"`
}
