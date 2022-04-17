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
	Role       string             `bson:"role" json:"role"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

func NewUser(id string, name string, login string, pass string, Role string) (user User, err error) {
	if id == "" {
		user.ID = primitive.NewObjectID()
	} else {
		user.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	user.Name = name
	user.Login = login
	user.Pass = pass
	user.Role = Role
	return
}

type UserInfo struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name" json:"name"`
	Login      string             `bson:"login" json:"login"`
	Role       string             `bson:"role" json:"role"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}
