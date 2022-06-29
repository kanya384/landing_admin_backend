package settings

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION = "settings"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Settings {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context) (settings []*domain.Setting, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{}, options.Find())
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var setting *domain.Setting
		err = cur.Decode(&setting)
		if err != nil {
			return
		}

		settings = append(settings, setting)
	}

	if err = cur.Err(); err != nil {
		return
	}

	cur.Close(ctx)
	return
}

func (r *repository) GetByID(ctx context.Context, settingID primitive.ObjectID) (setting *domain.Setting, err error) {
	err = r.collection.FindOne(ctx, primitive.M{"_id": settingID}).Decode(&setting)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return setting, errors.New(domain.ErrNoFieldWithID)
	}
	return
}

func (r *repository) GetByName(ctx context.Context, name string) (setting *domain.Setting, err error) {
	err = r.collection.FindOne(ctx, primitive.M{"name": name}).Decode(&setting)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return setting, errors.New(domain.ErrNoFieldWithID)
	}
	return
}

func (r *repository) Create(ctx context.Context, setting domain.Setting) (err error) {
	_, err = r.collection.InsertOne(ctx, &setting)
	return
}

func (r *repository) Update(ctx context.Context, setting domain.Setting) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": setting.ID}, primitive.D{{Key: "$set", Value: primitive.M{"name": setting.Name, "value": setting.Value}}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
