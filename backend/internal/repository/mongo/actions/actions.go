package actions

import (
	"context"
	"errors"
	"fmt"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION = "actions"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Action {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context, filter map[string]interface{}) (actions []*domain.Action, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{}, options.Find().SetSort(primitive.D{{"order", 1}, {"updated_at", -1}}))
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var action *domain.Action
		err = cur.Decode(&action)
		if err != nil {
			return
		}

		actions = append(actions, action)
	}

	if err = cur.Err(); err != nil {
		return
	}

	if len(actions) == 0 {
		return actions, mongo.ErrNoDocuments
	}

	cur.Close(ctx)
	return
}

func (r *repository) GetByID(ctx context.Context, id primitive.ObjectID) (action domain.Action, err error) {
	err = r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&action)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return action, errors.New(domain.ErrNoFieldWithID)
	}
	return
}
func (r *repository) Create(ctx context.Context, action domain.Action) (err error) {
	_, err = r.collection.InsertOne(ctx, &action)
	return
}
func (r *repository) Update(ctx context.Context, action domain.Action) (err error) {
	fmt.Println(action.Preview)
	fmt.Println(action.Photo)
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": action.ID}, primitive.D{{Key: "$set", Value: primitive.M{"title": action.Title, "description": action.Description, "date": action.Date, "photo": action.Photo, "preview": action.Preview, "order": action.Order, "updated_at": action.UpdateAt, "modified_by": action.ModifiedBy}}})
	fmt.Println(err)
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	case result.MatchedCount > 0 && result.ModifiedCount == 0:
		return errors.New(domain.ErrNoModificationsForEntity)
	}
	return
}

func (r *repository) UpdateOrder(ctx context.Context, id primitive.ObjectID, order int) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": id}, primitive.D{{Key: "$set", Value: primitive.M{"order": order}}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}

func (r *repository) Delete(ctx context.Context, actionID primitive.ObjectID) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": actionID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
