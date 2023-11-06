package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserModel struct {
	ID               primitive.ObjectID `bson:"_id"`
	UserName         string             `json:"userName" bson:"userName" validate:"required"`
	FirstName        string             `json:"firstName" bson:"firstName" validate:"required"`
	LastName         string             `json:"lastName" bson:"lastName" validate:"required"`
	ProfilePic       string             `json:"profilePic" bson:"profilePic" validate:"required"`
	Roles            []string           `json:"roles" bson:"roles" validate:"required"`
	IsActive         bool               `json:"isActive" bson:"isActive"`
	Status           string             `json:"status" bson:"status"`
	VerificationCode string             `json:"verificationCode" bson:"verificationCode"`
	IsVerified       bool               `json:"isVerified" bson:"isVerified"`
	Password         string             `json:"password" bson:"password" validate:"required"`
	Email            string             `json:"email" bson:"email" validate:"required"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt" bson:"updatedAt"`
}
