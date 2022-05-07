package hod_photos

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION = "photos"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.HodPhotos {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context, monthID primitive.ObjectID) (hodPhotos []domain.HodPhoto, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{"month_id": monthID}, options.Find().SetSort(primitive.D{{"order", 1}, {"updated_at", -1}}))
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var photo domain.HodPhoto
		err = cur.Decode(&photo)
		if err != nil {
			return
		}

		hodPhotos = append(hodPhotos, photo)
	}

	if err = cur.Err(); err != nil {
		return
	}

	if len(hodPhotos) == 0 {
		return hodPhotos, mongo.ErrNoDocuments
	}

	cur.Close(ctx)

	return
}

func (r *repository) Create(ctx context.Context, hodPhoto domain.HodPhoto) (err error) {
	_, err = r.collection.InsertOne(ctx, &hodPhoto)
	return
}

func (r *repository) Delete(ctx context.Context, hodPhotoID primitive.ObjectID) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": hodPhotoID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
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
