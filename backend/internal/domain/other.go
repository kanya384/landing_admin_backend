package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

//update order info
type UpdateOrder struct {
	ID    primitive.ObjectID `bson:"id" json:"id" form:"id" binding:"required"`
	Order int                `bson:"order" json:"order" form:"order" binding:"required"`
}

func NewUpdateOrderStruct(id string, order int) (updateOrder UpdateOrder, err error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	updateOrder = UpdateOrder{
		ID:    ID,
		Order: order,
	}
	return
}
