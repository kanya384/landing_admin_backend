package posters

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION = "posters"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Poster {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context, filter map[string]interface{}) (posters []*domain.Poster, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{})
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var poster *domain.Poster
		err = cur.Decode(&poster)
		if err != nil {
			return
		}

		posters = append(posters, poster)
	}

	if err = cur.Err(); err != nil {
		return
	}

	if len(posters) == 0 {
		return posters, mongo.ErrNoDocuments
	}

	cur.Close(ctx)

	return
}

func (r *repository) GetByID(ctx context.Context, id primitive.ObjectID) (poster domain.Poster, err error) {
	err = r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&poster)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return poster, errors.New(domain.ErrNoFieldWithID)
	}
	return
}
func (r *repository) Create(ctx context.Context, poster domain.Poster) (err error) {
	_, err = r.collection.InsertOne(ctx, &poster)
	return
}
func (r *repository) Update(ctx context.Context, poster domain.Poster) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": poster.ID}, primitive.D{{Key: "$set", Value: primitive.M{"title": poster.Title, "description": poster.Description, "photo": poster.Photo, "active": poster.Active, "order": poster.Order, "updated_at": poster.UpdateAt, "modified_by": poster.ModifiedBy}}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	case result.MatchedCount > 0 && result.ModifiedCount == 0:
		return errors.New(domain.ErrNoModificationsForEntity)
	}
	return
}

func (r *repository) Delete(ctx context.Context, posterID primitive.ObjectID) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": posterID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
