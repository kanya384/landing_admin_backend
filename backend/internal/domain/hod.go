package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Year struct {
	ID         primitive.ObjectID `bson:"_id"`
	Value      int                `bson:"value" json:"value" form:"value" binding:"required,numeric"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

type Month struct {
	ID         primitive.ObjectID `bson:"_id"`
	YearID     primitive.ObjectID `bson:"year_id" json:"year_id"`
	Value      int                `bson:"value" json:"value" form:"value" binding:"required,numeric,min=1,max=12"`
	Name       string             `bson:"name" json:"name"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

type HodPhoto struct {
	ID         primitive.ObjectID `bson:"_id"`
	MonthID    primitive.ObjectID `bson:"month_id" json:"month_id"`
	Image      string             `bson:"image" json:"image"`
	Order      int                `bson:"order" json:"order"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}
