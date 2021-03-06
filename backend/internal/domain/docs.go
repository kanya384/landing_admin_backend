package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//доки
type Doc struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"name" json:"title" form:"title" binding:"required"`
	File       string             `bson:"file" json:"file" form:"file" binding:"required"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

func NewDoc(id string, title string, file string, CreatedAt time.Time, UpdateAt time.Time, ModifiedBy primitive.ObjectID) (doc Doc, err error) {
	doc = Doc{
		Title:      title,
		File:       file,
		CreatedAt:  CreatedAt,
		UpdateAt:   UpdateAt,
		ModifiedBy: ModifiedBy,
	}
	if id == "" {
		doc.ID = primitive.NewObjectID()
	} else {
		doc.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}
