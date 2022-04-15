package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Plan struct {
	ID                  primitive.ObjectID `bson:"_id"`
	Entrance            int                `bson:"entrance" json:"entrance"`
	Commerce            bool               `bson:"commerce" json:"commerce"`
	Floor               int                `bson:"floor" json:"floor"`
	Number              int                `bson:"number" json:"number"`
	Rooms               int                `bson:"rooms" json:"rooms"`
	Area                float32            `bson:"area" json:"area"`
	LivingArea          float32            `bson:"living_area" json:"living_area"`
	KitchenArea         float32            `bson:"kitchen_area" json:"kitchen_area"`
	Price               int                `bson:"prcie" json:"prcie"`
	DiscountPrice       int                `bson:"discount_price" json:"discount_price"`
	SquarePrice         int                `bson:"square_price" json:"square_price"`
	SquareDiscountPrice int                `bson:"square_discount_price" json:"square_discount_price"`
	Discount            float32            `bson:"discount" json:"discount"`
	Status              bool               `bson:"status" json:"status"`
	WindowView          string             `bson:"window_view" json:"window_view"`
	Liter               string             `bson:"liter" json:"liter"`
	CreatedAt           time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt            time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy          primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}
