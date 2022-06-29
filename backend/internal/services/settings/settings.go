package settings

import (
	"context"
	"fmt"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/internal/services/pdf_generator"
)

type Settings interface {
	Get(ctx context.Context) (settings []*domain.Setting, err error)
	CreateOrUpdate(ctx context.Context, setting domain.Setting) (settingRes *domain.Setting, err error)
	GetByName(ctx context.Context, name string) (setting *domain.Setting, err error)
}

type service struct {
	repository   *repository.Repository
	pdfGenerator pdf_generator.PdfGenerator
	cfg          *config.Config
}

func NewService(repository *repository.Repository, cfg *config.Config, pdfGenerator pdf_generator.PdfGenerator) Settings {
	return &service{
		repository:   repository,
		pdfGenerator: pdfGenerator,
		cfg:          cfg,
	}
}

func (s *service) Get(ctx context.Context) (settings []*domain.Setting, err error) {
	return s.repository.Settings.Get(ctx)
}

func (s *service) GetByName(ctx context.Context, name string) (setting *domain.Setting, err error) {
	return s.repository.Settings.GetByName(ctx, name)
}

func (s *service) CreateOrUpdate(ctx context.Context, setting domain.Setting) (settingRes *domain.Setting, err error) {
	sett, err := s.repository.Settings.GetByName(ctx, setting.Name)
	settingRes = &setting
	if err != nil {
		err = s.repository.Settings.Create(ctx, setting)
		return
	}
	setting.ID = sett.ID
	err = s.repository.Settings.Update(ctx, setting)

	if err == nil && setting.Name == "remont_square_price" {
		err = s.pdfGenerator.GeneratePdfs(ctx, setting.Value)
		fmt.Println(err)
	}
	return settingRes, err
}
