package mongo

import (
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/internal/repository/mongo/advantage_photo"
	"landing_admin_backend/internal/repository/mongo/advantages"
	"landing_admin_backend/internal/repository/mongo/docs"
	"landing_admin_backend/internal/repository/mongo/editable"
	"landing_admin_backend/internal/repository/mongo/hod_photos"
	"landing_admin_backend/internal/repository/mongo/leads"
	"landing_admin_backend/internal/repository/mongo/months"
	"landing_admin_backend/internal/repository/mongo/plans"
	"landing_admin_backend/internal/repository/mongo/posters"
	"landing_admin_backend/internal/repository/mongo/project_info"
	"landing_admin_backend/internal/repository/mongo/users"
	"landing_admin_backend/internal/repository/mongo/video"
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
		Plans:          plans.NewRepository(db),
		Video:          video.NewRepository(db),
		Docs:           docs.NewRepository(db),
		Leads:          leads.NewRepository(db),
		Editable:       editable.NewRepository(db),
		ProjectInfo:    project_info.NewRepository(db),
	}
}
