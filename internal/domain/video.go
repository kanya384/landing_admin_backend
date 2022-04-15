package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//блок с видео
type Video struct {
	ID         primitive.ObjectID `bson:"_id"`
	URL        string             `bson:"url" json:"url" form:"url"`
	Preview    string             `bson:"preview" json:"preview" form:"preview"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}
