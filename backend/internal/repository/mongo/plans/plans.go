package plans

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION = "plans"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Plans {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context, filter map[string]interface{}) (plans []*domain.Plan, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{}, options.Find().SetSort(primitive.D{{"rooms", 1}}))
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var plan *domain.Plan
		err = cur.Decode(&plan)
		if err != nil {
			return
		}

		plans = append(plans, plan)
	}

	if err = cur.Err(); err != nil {
		return
	}

	cur.Close(ctx)
	return
}

func (r *repository) GetByID(ctx context.Context, id string) (plan domain.Plan, err error) {
	err = r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&plan)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return plan, errors.New(domain.ErrNoFieldWithID)
	}
	return
}

func (r *repository) Create(ctx context.Context, plan domain.Plan) (err error) {
	_, err = r.collection.InsertOne(ctx, &plan)
	return
}

func (r *repository) Update(ctx context.Context, plan domain.Plan) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": plan.ID}, primitive.D{{Key: "$set", Value: plan}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}

func (r *repository) Delete(ctx context.Context, planID string) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": planID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
