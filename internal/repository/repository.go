package repository

import (
	"database/sql"

	"github.com/yendefrr/more-work/internal/models"

	"github.com/yendefrr/more-work/internal/repository/mysqldb"
)

type Jobs interface {
	GetAll() ([]models.Job, error)
	GetByEmployer(employer string) ([]models.Job, error)
}

type Employers interface {
	GetAll() ([]models.Employer, error)
	GetByID(id int) (models.Employer, error)
}

type Seekers interface {
	GetAll() ([]models.Seeker, error)
	GetByID(id int) (models.Seeker, error)
}

type Repositories struct {
	Jobs      Jobs
	Employers Employers
	Seekers   Seekers
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Jobs:      mysqldb.NewJobsRepo(db),
		Employers: mysqldb.NewEmployersRepo(db),
		Seekers:   mysqldb.NewSeekersRepo(db),
	}
}
