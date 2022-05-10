package video

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION = "videos"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Video {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context) (videos []*domain.Video, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{}, options.Find().SetSort(primitive.D{{"status", -1}, {"rooms", 1}}))
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var video *domain.Video
		err = cur.Decode(&video)
		if err != nil {
			return
		}

		videos = append(videos, video)
	}

	if err = cur.Err(); err != nil {
		return
	}

	cur.Close(ctx)
	return
}
func (r *repository) Create(ctx context.Context, video domain.Video) (err error) {
	_, err = r.collection.InsertOne(ctx, &video)
	return
}

func (r *repository) Update(ctx context.Context, video domain.Video) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": video.ID}, primitive.D{{Key: "$set", Value: primitive.M{"url": video.URL, "preview": video.Preview, "updated_at": video.UpdateAt, "modified_by": video.ModifiedBy}}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}

func (r *repository) Delete(ctx context.Context, videoID primitive.ObjectID) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": videoID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
