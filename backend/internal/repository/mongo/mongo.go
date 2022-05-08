package mongo

import (
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/internal/repository/mongo/advantage_photo"
	"landing_admin_backend/internal/repository/mongo/advantages"
	"landing_admin_backend/internal/repository/mongo/hod_photos"
	"landing_admin_backend/internal/repository/mongo/months"
	"landing_admin_backend/internal/repository/mongo/posters"
	"landing_admin_backend/internal/repository/mongo/users"
	"landing_admin_backend/internal/repository/mongo/years"

	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(db *mongo.Database) *repository.Repository {
	return &repository.Repository{
		Users:          users.NewRepository(db),
		Poster:         posters.NewRepository(db),
		Year:           years.NewRepository(db),
		Month:          months.NewRepository(db),
		HodPhotos:      hod_photos.NewRepository(db),
		Advantages:     advantages.NewRepository(db),
		AdvantagePhoto: advantage_photo.NewRepository(db),
	}
}
