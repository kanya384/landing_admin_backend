package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Setting struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" josn:"description"`
	Value       int                `bson:"value" json:"value"`
}

func NewSetting(id string, name string, description string, value int) (setting Setting, err error) {
	setting = Setting{
		Name:        name,
		Value:       value,
		Description: description,
	}
	if id == "" {
		setting.ID = primitive.NewObjectID()
	} else {
		setting.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}
