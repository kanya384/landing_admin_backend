package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//контент для блока о проекте
type ProjectInfo struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title" json:"title"`
	Anonce      string             `bson:"anonce" json:"anonce"`
	Description string             `bson:"description" json:"description"`
	Photo       string             `bson:"photo" json:"photo"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt    time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy  primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}
