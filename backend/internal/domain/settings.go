package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Setting struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name" json:"name"`
	Value interface{}        `bson:"value" json:"value"`
}

func NewSetting(id string, name string, value interface{}) (setting Setting, err error) {
	setting = Setting{
		Name:  name,
		Value: value,
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
