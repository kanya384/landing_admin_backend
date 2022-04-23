package repository

import (
	"context"
	"landing_admin_backend/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users interface {
	Get(ctx context.Context) (users []*domain.UserInfo, err error)
	GetUserByCredetinals(ctx context.Context, login string, pass string) (user domain.UserInfo, err error)
	Create(ctx context.Context, user domain.User) (err error)
	Update(ctx context.Context, user domain.User) (err error)
	Delete(ctx context.Context, userID primitive.ObjectID) (err error)
}

type Advantages interface {
	Get(ctx context.Context) (advantages []domain.Advantage, err error)
	Create(ctx context.Context, advantage domain.Advantage) (err error)
	Update(ctx context.Context, advantage domain.Advantage) (err error)
	Delete(ctx context.Context, advantageID primitive.ObjectID) (err error)
}

type AdvantagePhoto interface {
	Create(ctx context.Context, advantagePhoto domain.AdvantagePhoto) (err error)
	Delete(ctx context.Context, advantagePhotoID primitive.ObjectID) (err error)
}

type Docs interface {
	Get(ctx context.Context) (docs []domain.Doc, err error)
	Create(ctx context.Context, doc domain.Doc) (err error)
	Update(ctx context.Context, doc domain.Doc) (err error)
	Delete(ctx context.Context, docID primitive.ObjectID) (err error)
}

type Year interface {
	Get(ctx context.Context) (docs []domain.Year, err error)
	Create(ctx context.Context, year domain.Year) (err error)
	Update(ctx context.Context, year domain.Year) (err error)
	Delete(ctx context.Context, yearID primitive.ObjectID) (err error)
}

type Month interface {
	Get(ctx context.Context) (months []domain.Month, err error)
	Create(ctx context.Context, month domain.Month) (err error)
	Update(ctx context.Context, month domain.Month) (err error)
	Delete(ctx context.Context, monthID primitive.ObjectID) (err error)
}

type HodPhotos interface {
	Get(ctx context.Context, monthID primitive.ObjectID) (hodPhotos []domain.HodPhoto, err error)
	Create(ctx context.Context, hodPhoto domain.HodPhoto) (err error)
	Delete(ctx context.Context, hodPhotoID primitive.ObjectID) (err error)
}

type Plans interface {
	CreateOrUpdate(ctx context.Context, plan domain.Plan) (err error)
	UpdateImage(ctx context.Context, id primitive.ObjectID, image string) (err error)
}

type Poster interface {
	Get(ctx context.Context, filter map[string]interface{}) (posters []*domain.Poster, err error)
	GetByID(ctx context.Context, id primitive.ObjectID) (poster domain.Poster, err error)
	Create(ctx context.Context, poster domain.Poster) (err error)
	Update(ctx context.Context, poster domain.Poster) (err error)
	UpdateOrder(ctx context.Context, id primitive.ObjectID, order int) (err error)
	Delete(ctx context.Context, posterID primitive.ObjectID) (err error)
}

type ProjectInfo interface {
	Get() (projectInfoList []domain.ProjectInfo, err error)
	Create(projectInfo domain.ProjectInfo) (err error)
	Update(projectInfo domain.ProjectInfo) (err error)
	Delete(projectInfoID primitive.ObjectID) (err error)
}

type Video interface {
	Get() (video []domain.Video, err error)
	Create(video domain.Video) (err error)
	Update(video domain.Video) (err error)
	Delete(videoID primitive.ObjectID) (err error)
}
