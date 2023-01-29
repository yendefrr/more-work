package mysqldb

import (
	"database/sql"

	"github.com/yendefrr/more-work/internal/models"
)

type EmployersRepo struct {
	db *sql.DB
}

func NewEmployersRepo(db *sql.DB) *EmployersRepo {
	return &EmployersRepo{
		db: db,
	}
}

func (r *EmployersRepo) GetAll() ([]models.Employer, error) {

	return nil, nil
}

func (r *EmployersRepo) GetByID(id int) (models.Employer, error) {
	var employer models.Employer

	return employer, nil
}
