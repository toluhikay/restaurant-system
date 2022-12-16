package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `json:"_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	Food_image *string            `json:"food_image" validate:"required"`
	Price      *float64           `json:"price" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Menu_id    *string            `json:"menu_id" validate:"required"`
	Food_id    string             `json:"food_id"`
}
