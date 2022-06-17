package project_info

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"

	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION = "pojectInfo"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.ProjectInfo {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context, filter map[string]interface{}) (pojectInfo []*domain.ProjectInfo, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{}, options.Find().SetSort(primitive.D{{"order", 1}, {"updated_at", -1}}))
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var projectInfo *domain.ProjectInfo
		err = cur.Decode(&projectInfo)
		if err != nil {
			return
		}

		pojectInfo = append(pojectInfo, projectInfo)
	}

	if err = cur.Err(); err != nil {
		return
	}

	if len(pojectInfo) == 0 {
		return pojectInfo, mongo.ErrNoDocuments
	}

	cur.Close(ctx)

	return
}

func (r *repository) GetByID(ctx context.Context, id primitive.ObjectID) (projectInfo domain.ProjectInfo, err error) {
	err = r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&projectInfo)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return projectInfo, errors.New(domain.ErrNoFieldWithID)
	}
	return
}
func (r *repository) Create(ctx context.Context, projectInfo domain.ProjectInfo) (err error) {
	_, err = r.collection.InsertOne(ctx, &projectInfo)
	return
}
func (r *repository) Update(ctx context.Context, projectInfo domain.ProjectInfo) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": projectInfo.ID}, primitive.D{{Key: "$set", Value: primitive.M{"title": projectInfo.Title, "description": projectInfo.Description, "photo": projectInfo.Photo, "anonce": projectInfo.Anonce, "order": projectInfo.Order, "updated_at": projectInfo.UpdateAt, "modified_by": projectInfo.ModifiedBy}}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	case result.MatchedCount > 0 && result.ModifiedCount == 0:
		return errors.New(domain.ErrNoModificationsForEntity)
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

func (r *repository) Delete(ctx context.Context, projectInfoID primitive.ObjectID) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": projectInfoID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
