package service

import (
	"context"

	"github.com/yendefrr/more-work/internal/repository"
)

type SeekersService struct {
	repo repository.Seekers
}

func NewSeekersService(repo repository.Seekers) *SeekersService {
	return &SeekersService{
		repo: repo,
	}
}

func (s *SeekersService) Create(ctx context.Context, form SeekerCreateForm) error {

	return nil
}
