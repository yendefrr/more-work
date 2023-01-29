package mysqldb

import (
	"database/sql"

	"github.com/yendefrr/more-work/internal/models"
)

type JobsRepo struct {
	db *sql.DB
}

func NewJobsRepo(db *sql.DB) *JobsRepo {
	return &JobsRepo{
		db: db,
	}
}

func (r *JobsRepo) GetAll() ([]models.Job, error) {

	return nil, nil
}

func (r *JobsRepo) GetByEmployer(employer string) ([]models.Job, error) {

	return nil, nil
}
