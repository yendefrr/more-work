package service

import (
	"context"

	"github.com/yendefrr/more-work/internal/repository"
)

type JobsService struct {
	repo repository.Jobs
}

func NewJobsService(repo repository.Jobs) *JobsService {
	return &JobsService{
		repo: repo,
	}
}

func (s *JobsService) Create(ctx context.Context, form JobCreateForm) error {

	return nil
}
