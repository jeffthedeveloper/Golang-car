package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TCar struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Model    string             `json:"model,omitempty" bson:"model,omitempty" validate:"required"`
	Year     int                `json:"year,omitempty" bson:"year,omitempty" validate:"required"`
	Color    string             `json:"color,omitempty" bson:"color,omitempty" validate:"required"`
	Status   bool               `json:"status,omitempty" bson:"status,omitempty" validate:"required"`
	BuyValue int                `json:"buyValue,omitempty" bson:"buyValue,omitempty" validate:"required"`
	DoorsQty int                `json:"doorsQty,omitempty" bson:"doorsQty,omitempty" validate:"required,gte=2,lte=4"`
	SeatsQty int                `json:"seatsQty,omitempty" bson:"seatsQty,omitempty" validate:"required,gte=2,lte=5"`
}
