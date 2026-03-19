package sermon

import (
	"church-backend/internal/models"
	"context"
)

type SermonRepository interface {
	GetLatest(ctx context.Context, limit int) ([]models.Sermon, error)
	CreateSermon(ctx context.Context, s *models.Sermon) error
}

type Service struct {
	repo SermonRepository
}

func NewService(repo SermonRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetRecent(ctx context.Context, limit int) ([]models.Sermon, error) {
	if limit <= 0 {
		limit = 5
	}
	return s.repo.GetLatest(ctx, limit)
}
