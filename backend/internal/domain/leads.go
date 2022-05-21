package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lead struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name" json:"name" form:"name"`
	Phone        string             `bson:"phone" json:"phone" form:"phone"`
	Email        string             `bson:"email" json:"email" form:"email"`
	Text         string             `bson:"text" json:"text" form:"text"`
	Roistat      string             `bson:"roistat" json:"roistat" form:"roistat"`
	Utm_source   string             `bson:"utm_source" json:"utm_source" form:"utm_source"`
	Utm_medium   string             `bson:"utm_medium" json:"utm_medium" form:"utm_medium"`
	Utm_term     string             `bson:"utm_term" json:"utm_term" form:"utm_term"`
	Utm_campaign string             `bson:"utm_campaign" json:"utm_campaign" form:"utm_campaign"`
	Utm_content  string             `bson:"utm_content" json:"utm_content" form:"utm_content"`
	SendedToCrm  bool               `bson:"sended_to_crm" json:"sended_to_crm" form:"sended_to_crm"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	ModifiedBy   primitive.ObjectID `bson:"modified_by" json:"modified_by"`
}

func NewLead(id string, name string, phone string, email string, text string, roistat string, utm_source string, utm_medium string, utm_term string, utm_campaign string, utm_content string, CreatedAt time.Time, UpdatedAt time.Time, ModifiedBy primitive.ObjectID) (lead Lead, err error) {
	lead = Lead{
		Name:         name,
		Phone:        phone,
		Email:        email,
		Text:         text,
		Roistat:      roistat,
		Utm_source:   utm_source,
		Utm_medium:   utm_medium,
		Utm_term:     utm_term,
		Utm_campaign: utm_campaign,
		Utm_content:  utm_content,
		SendedToCrm:  false,
		CreatedAt:    CreatedAt,
		UpdatedAt:    UpdatedAt,
		ModifiedBy:   ModifiedBy,
	}
	if id == "" {
		lead.ID = primitive.NewObjectID()
	} else {
		lead.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return
		}
	}
	return
}
