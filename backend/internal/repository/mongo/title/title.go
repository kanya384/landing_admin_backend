package title

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION = "titles"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Title {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context) (titles []*domain.Title, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{})
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var title domain.Title
		err = cur.Decode(&title)
		if err != nil {
			return
		}

		titles = append(titles, &title)
	}

	if err = cur.Err(); err != nil {
		return
	}

	if len(titles) == 0 {
		return titles, mongo.ErrNoDocuments
	}

	cur.Close(ctx)

	return
}
func (r *repository) GetByTag(ctx context.Context, tagName string, tagValue string) (title *domain.Title, err error) {
	err = r.collection.FindOne(ctx, primitive.M{"tag_name": tagName, "tag_value": tagValue}).Decode(&title)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return title, errors.New(domain.ErrCheckUserAndPass)
	}
	return
}
func (r *repository) Create(ctx context.Context, title domain.Title) (err error) {
	_, err = r.collection.InsertOne(ctx, &title)
	return
}
func (r *repository) Update(ctx context.Context, title domain.Title) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": title.ID}, primitive.D{{Key: "$set", Value: title}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	case result.MatchedCount > 0 && result.ModifiedCount == 0:
		return errors.New(domain.ErrNoModificationsForEntity)
	}
	return
}
func (r *repository) Delete(ctx context.Context, titleID primitive.ObjectID) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": titleID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
