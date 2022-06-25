package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Title struct {
	ID           primitive.ObjectID `bson:"_id"`
	TagName      string             `bson:"tag_name" json:"tag_name"`   //метка по которой подменяется тайтл
	TagValue     string             `bson:"tag_value" json:"tag_value"` //значение метки по которой подменяется тайтл
	DesktopTitle string             `bson:"desktop_title" json:"desktop_title"`
	MobileTitle  string             `bson:"mobile_title" json:"mobile_title"`
}
