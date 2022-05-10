package docs

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION = "docs"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Docs {
	collection := db.Collection(COLLECTION)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context) (docs []*domain.Doc, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{}, options.Find().SetSort(primitive.D{{"created_at", -1}}))
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var doc *domain.Doc
		err = cur.Decode(&doc)
		if err != nil {
			return
		}

		docs = append(docs, doc)
	}

	if err = cur.Err(); err != nil {
		return
	}

	cur.Close(ctx)
	return
}
func (r *repository) Create(ctx context.Context, doc domain.Doc) (err error) {
	_, err = r.collection.InsertOne(ctx, &doc)
	return
}

func (r *repository) Update(ctx context.Context, doc domain.Doc) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": doc.ID}, primitive.D{{Key: "$set", Value: primitive.M{"title": doc.Title, "file": doc.File, "updated_at": doc.UpdateAt, "modified_by": doc.ModifiedBy}}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}

func (r *repository) Delete(ctx context.Context, docID primitive.ObjectID) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": docID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
