package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Year struct {
	ID         primitive.ObjectID `bson:"_id"`
	Value      int                `bson:"value" json:"value" form:"value" binding:"required,numeric"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

func NewYear(id string, value int, createdAt time.Time, updateAt time.Time, modifiedBy primitive.ObjectID) (year Year, err error) {
	year = Year{
		Value:      value,
		CreatedAt:  createdAt,
		UpdateAt:   updateAt,
		ModifiedBy: modifiedBy,
	}
	if id == "" {
		year.ID = primitive.NewObjectID()
	} else {
		year.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}

type Month struct {
	ID         primitive.ObjectID `bson:"_id"`
	YearID     primitive.ObjectID `bson:"year_id" json:"year_id"`
	Value      int                `bson:"value" json:"value" form:"value" binding:"required,numeric,min=1,max=12"`
	Name       string             `bson:"name" json:"name"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

func NewMonth(id string, yearID string, value int, name string, createdAt time.Time, updateAt time.Time, modifiedBy primitive.ObjectID) (month Month, err error) {
	month = Month{
		Value:      value,
		CreatedAt:  createdAt,
		UpdateAt:   updateAt,
		ModifiedBy: modifiedBy,
	}
	month.Name = getMonthNamFromInt(value)

	month.YearID, err = primitive.ObjectIDFromHex(yearID)
	if err != nil {
		return
	}

	if id == "" {
		month.ID = primitive.NewObjectID()
	} else {
		month.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}

type HodPhoto struct {
	ID         primitive.ObjectID `bson:"_id"`
	MonthID    primitive.ObjectID `bson:"month_id" json:"month_id"`
	Image      string             `bson:"image" json:"image"`
	Order      int                `bson:"order" json:"order"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt   time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

func NewHodPhoto(id string, monthID string, image string, order int, createdAt time.Time, updateAt time.Time, modifiedBy primitive.ObjectID) (photo HodPhoto, err error) {
	photo = HodPhoto{
		Image:      image,
		Order:      order,
		CreatedAt:  createdAt,
		UpdateAt:   updateAt,
		ModifiedBy: modifiedBy,
	}

	photo.MonthID, err = primitive.ObjectIDFromHex(monthID)
	if err != nil {
		return
	}
	if id == "" {
		photo.ID = primitive.NewObjectID()
	} else {
		photo.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}

func getMonthNamFromInt(value int) string {
	months := map[int]string{
		0:  "Январь",
		1:  "Февраль",
		2:  "Март",
		3:  "Апрель",
		4:  "Май",
		5:  "Июнь",
		6:  "Июль",
		7:  "Август",
		8:  "Сентябрь",
		9:  "Октябрь",
		10: "Ноябрь",
		11: "Декабрь",
	}
	return months[value]
}
