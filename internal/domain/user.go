package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name" json:"name"`
	Login      string             `bson:"login" json:"login"`
	Pass       string             `bson:"pass" json:"pass"`
	Role       int                `bson:"role" json:"role"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

type UserInfo struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name" json:"name"`
	Login      string             `bson:"login" json:"login"`
	Role       int                `bson:"role" json:"role"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}
