package domain

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Editable struct {
	ID     primitive.ObjectID `bson:"_id"`
	Type   int                `bson:"type" json:"type" form:"type" binding:"required"` //0 - text 1 - photo
	Value  string             `bson:"value" json:"value" form:"value" binding:"required"`
	Width  int                `bson:"width" json:"width" form:"width"`
	Height int                `bson:"height" json:"height" form:"height"`
}

func NewEditable(id string, Type int, value string, width int, height int) (editable Editable, err error) {
	fmt.Println(id)
	editable = Editable{
		Type:   Type,
		Value:  value,
		Width:  width,
		Height: height,
	}
	if id == "" {
		editable.ID = primitive.NewObjectID()
	} else {
		editable.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}
