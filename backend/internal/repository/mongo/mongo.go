package mongo

import (
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/internal/repository/mongo/users"

	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(db *mongo.Database) *repository.Repository {
	return &repository.Repository{
		Users: users.NewRepository(db),
	}
}
