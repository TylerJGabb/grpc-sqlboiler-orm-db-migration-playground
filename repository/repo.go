package repository

import (
	"database/sql"
	"sqlboiler-sb/models"
)

type ChangeRequestRepository interface {
	CreateChangeRequest(cr *models.ChangeRequest) error
	GetChangeRequest(id int) (*models.ChangeRequest, error)
	GetAllChangeRequests() ([]*models.ChangeRequest, error)
	UpdateChangeRequest(cr *models.ChangeRequest) error
	AddStatusEvents(cr *models.ChangeRequest, se *models.StatusEvent) error
}

type PostgresChangeRequestRepository struct {
	db *sql.DB
}

func NewPostgresChangeRequestRepository(
	dbname, user, password string,
) *PostgresChangeRequestRepository {
	return nil
}

func (r *PostgresChangeRequestRepository) CreateChangeRequest(cr *models.ChangeRequest) error {
	return nil
}
