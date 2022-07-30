package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//блок а так же вас ждет
type Action struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Date        string             `bson:"date" json:"date"`
	Preview     string             `bson:"preview" json:"preview"`
	Photo       string             `bson:"photo" json:"photo"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt    time.Time          `bson:"updated_at" json:"updated_at"`
	Order       int                `bson:"order" json:"order"`
	ModifiedBy  primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

func NewAction(id string, title string, description string, date string, preview string, photo string, order int, CreatedAt time.Time, UpdateAt time.Time, ModifiedBy primitive.ObjectID) (action Action, err error) {
	action = Action{
		Title:       title,
		Description: description,
		Date:        date,
		Preview:     preview,
		Order:       order,
		CreatedAt:   CreatedAt,
		UpdateAt:    UpdateAt,
		ModifiedBy:  ModifiedBy,
	}
	if id == "" {
		action.ID = primitive.NewObjectID()
	} else {
		action.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}
