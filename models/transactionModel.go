package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Type int

const (
	Income Type = iota
	Expense
)

type Transaction struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`                           // Maps the MongoDB `_id` field
	UserName    string             `json:"username" validate:"required"`            // UserName is required
	Type        Type               `json:"type" validate:"required,oneof=0 1"`      // Type must be either 0 (Income) or 1 (Expense)
	Description string             `json:"description" validate:"required,max=255"` // Description is required, max length 255 characters
	Amount      int                `json:"amount" validate:"required,gt=0"`         // Amount is required and must be greater than 0
}
