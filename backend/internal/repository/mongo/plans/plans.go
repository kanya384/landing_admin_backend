package plans

import (
	"context"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/mongo"
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

func (r *repository) Get(ctx context.Context, filter map[string]interface{}) (plans []domain.Plan, err error) {
	return
}
func (r *repository) Create(ctx context.Context, plan domain.Plan) (err error) {
	return
}
func (r *repository) Update(ctx context.Context, plan domain.Plan) (err error) {
	return
}
func (r *repository) Delete(ctx context.Context, planID string) (err error) {
	return
}
