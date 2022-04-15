package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//блок а так же вас ждет
type Advantage struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt    time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy  primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

type AdvantagePhoto struct {
	ID          primitive.ObjectID `bson:"_id"`
	AdvantageID primitive.ObjectID `bson:"advantage_id"`
	Image       string             `bson:"image" json:"image"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt    time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy  primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}
