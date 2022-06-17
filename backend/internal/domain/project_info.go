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
	Order       int                `bson:"order" json:"order"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt    time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy  primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

func NewProjectInfo(id string, title string, anonce string, description string, photo string, order int, CreatedAt time.Time, UpdateAt time.Time, ModifiedBy primitive.ObjectID) (projectInfo ProjectInfo, err error) {
	projectInfo = ProjectInfo{
		Title:       title,
		Anonce:      anonce,
		Description: description,
		Photo:       photo,
		Order:       order,
		CreatedAt:   CreatedAt,
		UpdateAt:    UpdateAt,
		ModifiedBy:  ModifiedBy,
	}
	if id == "" {
		projectInfo.ID = primitive.NewObjectID()
	} else {
		projectInfo.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}
