package months

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"
	"landing_admin_backend/internal/repository/mongo/hod_photos"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION = "months"

type repository struct {
	collection       *mongo.Collection
	photosCollection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Month {
	collection := db.Collection(COLLECTION)
	photosCollection := db.Collection(hod_photos.COLLECTION)
	return &repository{
		collection:       collection,
		photosCollection: photosCollection,
	}
}

func (r *repository) Get(ctx context.Context, yearID primitive.ObjectID) (months []domain.Month, err error) {
	var cur *mongo.Cursor
	if yearID == primitive.NilObjectID {
		cur, err = r.collection.Find(ctx, primitive.M{}, options.Find().SetSort(primitive.D{{"value", 1}}))
		if err != nil {
			return
		}
	} else {
		cur, err = r.collection.Find(ctx, primitive.M{"year_id": yearID}, options.Find().SetSort(primitive.D{{"value", 1}}))
		if err != nil {
			return
		}
	}
	for cur.Next(ctx) {
		var month domain.Month
		err = cur.Decode(&month)
		if err != nil {
			return
		}

		months = append(months, month)
	}

	if err = cur.Err(); err != nil {
		return
	}

	if len(months) == 0 {
		return months, mongo.ErrNoDocuments
	}

	cur.Close(ctx)

	return
}

func (r *repository) Create(ctx context.Context, month domain.Month) (err error) {
	_, err = r.collection.InsertOne(ctx, &month)
	return
}
func (r *repository) Update(ctx context.Context, month domain.Month) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": month.ID}, primitive.D{{Key: "$set", Value: primitive.M{"value": month.Value, "year_id": month.YearID, "name": month.Name, "updated_at": month.UpdateAt, "modified_by": month.ModifiedBy}}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	case result.MatchedCount > 0 && result.ModifiedCount == 0:
		return errors.New(domain.ErrNoModificationsForEntity)
	}
	return
}

func (r *repository) Delete(ctx context.Context, monthID primitive.ObjectID) (err error) {
	//удаляем фотки по месяцу
	_, err = r.photosCollection.DeleteMany(ctx, primitive.M{"month_id": monthID})
	if err != nil {
		return
	}

	//удаляем месяц
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": monthID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
