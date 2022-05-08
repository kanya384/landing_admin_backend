package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//блок а так же вас ждет
type Advantage struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt    time.Time          `bson:"updated_at" json:"updated_at"`
	Order       int                `bson:"order" json:"order"`
	ModifiedBy  primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

func NewAdvantage(id string, title string, description string, order int, CreatedAt time.Time, UpdateAt time.Time, ModifiedBy primitive.ObjectID) (advantage Advantage, err error) {
	advantage = Advantage{
		Title:       title,
		Description: description,
		Order:       order,
		CreatedAt:   CreatedAt,
		UpdateAt:    UpdateAt,
		ModifiedBy:  ModifiedBy,
	}
	if id == "" {
		advantage.ID = primitive.NewObjectID()
	} else {
		advantage.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}

type AdvantagePhoto struct {
	ID          primitive.ObjectID `bson:"_id"`
	AdvantageID primitive.ObjectID `bson:"advantage_id"`
	Image       string             `bson:"image" json:"image"`
	Order       int                `bson:"order" json:"order"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt    time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy  primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

func NewAdvantagePhoto(id string, advantageID string, image string, order int, CreatedAt time.Time, UpdateAt time.Time, ModifiedBy primitive.ObjectID) (advantagePhoto AdvantagePhoto, err error) {
	advantagePhoto = AdvantagePhoto{
		Image:      image,
		Order:      order,
		CreatedAt:  CreatedAt,
		UpdateAt:   UpdateAt,
		ModifiedBy: ModifiedBy,
	}
	advantagePhoto.AdvantageID, err = primitive.ObjectIDFromHex(advantageID)
	if err != nil {
		return
	}
	if id == "" {
		advantagePhoto.ID = primitive.NewObjectID()
	} else {
		advantagePhoto.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}
