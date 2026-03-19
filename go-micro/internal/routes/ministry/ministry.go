package ministry

import (
	"church-backend/internal/models"
	"context"
)

type MinistryRepository interface {
	GetAllMinistries(ctx context.Context) ([]models.Ministry, error)
	JoinMinistry(ctx context.Context, memberID string, ministryID string) error
}

type Service struct {
	repo MinistryRepository
}

func NewService(repo MinistryRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListMinistries(ctx context.Context) ([]models.Ministry, error) {
	return s.repo.GetAllMinistries(ctx)
}

func (s *Service) Join(ctx context.Context, phone, ministryID string) error {
	return s.repo.JoinMinistry(ctx, phone, ministryID)
}
