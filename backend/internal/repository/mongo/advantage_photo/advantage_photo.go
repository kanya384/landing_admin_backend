package advantage_photo

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION = "advantage_photo"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.AdvantagePhoto {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context, advantageID primitive.ObjectID) (advantagePhotos []domain.AdvantagePhoto, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{"month_id": advantageID}, options.Find().SetSort(primitive.D{{"order", 1}, {"updated_at", -1}}))
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var photo domain.AdvantagePhoto
		err = cur.Decode(&photo)
		if err != nil {
			return
		}

		advantagePhotos = append(advantagePhotos, photo)
	}

	if err = cur.Err(); err != nil {
		return
	}

	if len(advantagePhotos) == 0 {
		return advantagePhotos, mongo.ErrNoDocuments
	}

	cur.Close(ctx)

	return
}

func (r *repository) Create(ctx context.Context, advantagePhoto domain.AdvantagePhoto) (err error) {
	_, err = r.collection.InsertOne(ctx, &advantagePhoto)
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

func (r *repository) Delete(ctx context.Context, advantagePhotoID primitive.ObjectID) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": advantagePhotoID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
