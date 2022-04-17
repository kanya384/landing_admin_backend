package users

import (
	"context"
	"errors"
	"landing_admin_backend/internal/domain"
	repos "landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION = "users"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repos.Users {
	collection := db.Collection(COLLECTION)
	collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    primitive.D{{Key: "login", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	return &repository{
		collection: collection,
	}
}

func (r *repository) Get(ctx context.Context) (users []*domain.UserInfo, err error) {
	cur, err := r.collection.Find(ctx, primitive.M{})
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var user domain.UserInfo
		err = cur.Decode(&user)
		if err != nil {
			return
		}

		users = append(users, &user)
	}

	if err = cur.Err(); err != nil {
		return
	}

	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}

	cur.Close(ctx)

	return
}

func (r *repository) GetUserByCredetinals(ctx context.Context, login string, pass string) (user domain.UserInfo, err error) {
	err = r.collection.FindOne(ctx, primitive.M{"login": login, "pass": pass}).Decode(&user)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return user, errors.New(domain.ErrCheckUserAndPass)
	}
	return
}
func (r *repository) Create(ctx context.Context, user domain.User) (err error) {
	_, err = r.collection.InsertOne(ctx, &user)
	return
}
func (r *repository) Update(ctx context.Context, user domain.User) (err error) {
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": user.ID}, primitive.D{{Key: "$set", Value: user}})
	switch true {
	case result.MatchedCount == 0:
		return errors.New(domain.ErrNoFieldWithID)
	case result.MatchedCount > 0 && result.ModifiedCount == 0:
		return errors.New(domain.ErrNoModificationsForEntity)
	}
	return
}

func (r *repository) Delete(ctx context.Context, userID primitive.ObjectID) (err error) {
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": userID})
	if result.DeletedCount == 0 {
		return errors.New(domain.ErrNoFieldWithID)
	}
	return
}
