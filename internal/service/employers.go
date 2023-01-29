package service

import (
	"context"

	"github.com/yendefrr/more-work/internal/repository"
)

type EmployersService struct {
	repo repository.Employers
}

func NewEmployersService(repo repository.Employers) *EmployersService {
	return &EmployersService{
		repo: repo,
	}
}

func (s *EmployersService) Create(ctx context.Context, form EmployerCreateForm) error {

	return nil
}