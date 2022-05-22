package editable

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"

	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION = "editable"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Editable {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context) (editables []*domain.Editable, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{}, options.Find())
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var editable *domain.Editable
		err = cur.Decode(&editable)
		if err != nil {
			return
		}

		editables = append(editables, editable)
	}

	if err = cur.Err(); err != nil {
		return
	}

	cur.Close(ctx)
	return
}

func (r *repository) GetByID(ctx context.Context, id primitive.ObjectID) (editable domain.Editable, err error) {
	err = r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&editable)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return editable, errors.New(domain.ErrNoFieldWithID)
	}
	return
}

func (r *repository) Create(ctx context.Context, editable *domain.Editable) (err error) {
	_, err = r.collection.InsertOne(ctx, editable)
	return
}

func (r *repository) Update(ctx context.Context, editable *domain.Editable) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": editable.ID}, primitive.D{{Key: "$set", Value: editable}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
