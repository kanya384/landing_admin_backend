package advantages

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION = "advantages"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Advantages {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context) (advantages []*domain.Advantage, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{}, options.Find().SetSort(primitive.D{{"order", 1}, {"updated_at", -1}}))
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var advantage *domain.Advantage
		err = cur.Decode(&advantage)
		if err != nil {
			return
		}

		advantages = append(advantages, advantage)
	}

	if err = cur.Err(); err != nil {
		return
	}

	if len(advantages) == 0 {
		return advantages, mongo.ErrNoDocuments
	}

	cur.Close(ctx)

	return
}

func (r *repository) GetByID(ctx context.Context, id primitive.ObjectID) (advantage domain.Advantage, err error) {
	err = r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&advantage)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return advantage, errors.New(domain.ErrNoFieldWithID)
	}
	return
}

func (r *repository) Create(ctx context.Context, advantage domain.Advantage) (err error) {
	_, err = r.collection.InsertOne(ctx, &advantage)
	return
}

func (r *repository) Update(ctx context.Context, advantage domain.Advantage) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": advantage.ID}, primitive.D{{Key: "$set", Value: primitive.M{"title": advantage.Title, "description": advantage.Description, "order": advantage.Order, "updated_at": advantage.UpdateAt, "modified_by": advantage.ModifiedBy}}})
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

func (r *repository) Delete(ctx context.Context, advantageID primitive.ObjectID) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": advantageID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
