package mysqldb

import (
	"database/sql"

	"github.com/yendefrr/more-work/internal/models"
)

type SeekersRepo struct {
	db *sql.DB
}

func NewSeekersRepo(db *sql.DB) *SeekersRepo {
	return &SeekersRepo{
		db: db,
	}
}

func (r *SeekersRepo) GetAll() ([]models.Seeker, error) {

	return nil, nil
}

func (r *SeekersRepo) GetByID(id int) (models.Seeker, error) {
	var seeker models.Seeker

	return seeker, nil
}
