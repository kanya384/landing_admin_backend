package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//контент для первого блока
type Poster struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title" json:"title"` //нужен тайтл или нет?
	Photo      string             `bson:"photo" json:"photo"` //нужен тайтл или нет?
	Order      int                `bson:"order" json:"order"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

func NewPoster(id string, title string, photo string, order int, CreatedAt time.Time, UpdateAt time.Time, ModifiedBy primitive.ObjectID) (poster Poster, err error) {
	poster = Poster{
		Title:      title,
		Photo:      photo,
		Order:      order,
		CreatedAt:  CreatedAt,
		UpdateAt:   UpdateAt,
		ModifiedBy: ModifiedBy,
	}
	if id == "" {
		poster.ID = primitive.NewObjectID()
	} else {
		poster.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}
