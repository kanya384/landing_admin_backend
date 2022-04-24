package years

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"
	"landing_admin_backend/internal/repository/mongo/hod_photos"
	"landing_admin_backend/internal/repository/mongo/months"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION = "years"

type repository struct {
	collection       *mongo.Collection
	monthsCollection *mongo.Collection
	photosCollection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Year {
	collection := db.Collection(COLLECTION)
	collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    primitive.D{{Key: "value", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	monthsCollection := db.Collection(months.COLLECTION)
	photosCollection := db.Collection(hod_photos.COLLECTION)
	return &repository{
		collection:       collection,
		monthsCollection: monthsCollection,
		photosCollection: photosCollection,
	}
}

func (r *repository) Get(ctx context.Context) (years []domain.Year, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{}, options.Find().SetSort(primitive.D{{"value", -1}}))
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var year domain.Year
		err = cur.Decode(&year)
		if err != nil {
			return
		}

		years = append(years, year)
	}

	if err = cur.Err(); err != nil {
		return
	}

	if len(years) == 0 {
		return years, mongo.ErrNoDocuments
	}

	cur.Close(ctx)

	return
}

func (r *repository) Create(ctx context.Context, year domain.Year) (err error) {
	_, err = r.collection.InsertOne(ctx, &year)
	return
}
func (r *repository) Update(ctx context.Context, year domain.Year) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": year.ID}, primitive.D{{Key: "$set", Value: primitive.M{"value": year.Value, "updated_at": year.UpdateAt, "modified_by": year.ModifiedBy}}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	case result.MatchedCount > 0 && result.ModifiedCount == 0:
		return errors.New(domain.ErrNoModificationsForEntity)
	}
	return
}

func (r *repository) Delete(ctx context.Context, yearID primitive.ObjectID) (err error) {
	//находим все мясяца года
	months := []domain.Month{}
	cur, err := r.monthsCollection.Find(ctx, primitive.M{"year_id": yearID}, options.Find().SetSort(primitive.D{{"value", 1}}))
	if err != nil {
		return
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

	cur.Close(ctx)

	//удаляем фото по каждому месяцу
	for _, month := range months {
		_, err = r.photosCollection.DeleteMany(ctx, primitive.M{"month_id": month.ID})
		if err != nil {
			return
		}
		_, err = r.monthsCollection.DeleteOne(ctx, primitive.M{"_id": month.ID})
		if err != nil {
			return
		}
	}

	//удаляем год
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": yearID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
