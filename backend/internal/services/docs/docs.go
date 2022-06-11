package docs

import (
	"context"
	"io"
	"io/ioutil"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/pkg/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Docs interface {
	Get(ctx context.Context) (docs []*domain.Doc, err error)
	Create(ctx context.Context, doc domain.Doc, file io.ReadCloser, fileName string) (domain.Doc, error)
	Update(ctx context.Context, doc domain.Doc, file interface{}, fileName string) (domain.Doc, error)
	Delete(ctx context.Context, docID primitive.ObjectID) (err error)
}

type service struct {
	repository *repository.Repository
	cfg        *config.Config
}

func NewService(repository *repository.Repository, cfg *config.Config) Docs {
	return &service{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *service) Get(ctx context.Context) (docs []*domain.Doc, err error) {
	return s.repository.Docs.Get(ctx)
}

func (s *service) Create(ctx context.Context, doc domain.Doc, file io.ReadCloser, fileName string) (domain.Doc, error) {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return doc, err
	}
	name, err := helpers.SaveFile(fileBytes, fileName, s.cfg.FileStore)
	if err != nil {
		return doc, err
	}
	doc.File = name
	err = s.repository.Docs.Create(ctx, doc)
	return doc, err
}

func (s *service) Update(ctx context.Context, doc domain.Doc, file interface{}, fileName string) (domain.Doc, error) {
	if fileName != "" {
		fileBytes, err := ioutil.ReadAll(file.(io.ReadCloser))
		if err != nil {
			return doc, nil
		}
		name, err := helpers.SaveFile(fileBytes, fileName, s.cfg.FileStore)
		if err != nil {
			return doc, nil
		}
		doc.File = name
	}
	err := s.repository.Docs.Update(ctx, doc)
	return doc, err
}

func (s *service) Delete(ctx context.Context, docID primitive.ObjectID) (err error) {
	return s.repository.Docs.Delete(ctx, docID)
}
