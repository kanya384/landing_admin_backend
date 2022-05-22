package leads

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION = "leads"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Leads {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context) (leads []*domain.Lead, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{}, options.Find().SetLimit(50).SetSort(primitive.D{{"created_at", -1}}))
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var lead *domain.Lead
		err = cur.Decode(&lead)
		if err != nil {
			return
		}

		leads = append(leads, lead)
	}

	if err = cur.Err(); err != nil {
		return
	}

	cur.Close(ctx)
	return
}
func (r *repository) Create(ctx context.Context, lead domain.Lead) (err error) {
	_, err = r.collection.InsertOne(ctx, &lead)
	return
}

func (r *repository) Delete(ctx context.Context, leadID primitive.ObjectID) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": leadID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}

func (r *repository) CountDocuments(ctx context.Context, startDate time.Time, endDate time.Time) (count int, err error) {
	filter := primitive.D{{"created_at", primitive.D{{"$gt", startDate}}}, {"created_at", primitive.D{{"$lt", endDate}}}}
	countInt64, err := r.collection.CountDocuments(ctx, filter)
	return int(countInt64), err
}
