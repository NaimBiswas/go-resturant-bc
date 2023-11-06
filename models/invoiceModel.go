package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type InvoiceModel struct {
	ID             primitive.ObjectID `bson:"_id"`
	OrderId        string             `json:"orderId" validate:"required"`
	PaymentMethod  string             `json:"paymentMethod" validate:"eq=CARD|eq=CASH|eq="`
	PaymentStatus  string             `json:"paymentStatus" validate:"required"`
	PaymentDueDate string             `json:"paymentDueDate" validate:"required"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
}
