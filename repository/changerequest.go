package repository

import (
	"context"
	"database/sql"
	"sqlboiler-sb/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ChangeRequestRepository interface {
	CreateChangeRequest(cr *models.ChangeRequest) error
	GetChangeRequest(id int) (*models.ChangeRequest, error)
	GetAllChangeRequests() ([]*models.ChangeRequest, error)
	UpdateChangeRequest(cr *models.ChangeRequest) error
	AddTMTJob(cr *models.ChangeRequest, job *models.TMTJob) error
	AddRebaseJob(cr *models.ChangeRequest, job *models.RebaseJob) error
	UpdateTMTJob(job *models.TMTJob) error
	UpdateRebaseJob(job *models.RebaseJob) error
	GetTMTJob(id int) (*models.TMTJob, error)
	GetRebaseJob(id int) (*models.RebaseJob, error)
}

type SqlDbChangeRequestRepo struct {
	db *sql.DB
}

func (r *SqlDbChangeRequestRepo) CreateChangeRequest(cr *models.ChangeRequest) error {
	return cr.Insert(context.Background(), r.db, boil.Infer())
}

func (r *SqlDbChangeRequestRepo) GetChangeRequest(id int) (*models.ChangeRequest, error) {
	return models.ChangeRequests(
		qm.Load(models.ChangeRequestRels.TMTJobs),
		qm.Load(models.ChangeRequestRels.RebaseJobs),
		models.ChangeRequestWhere.ID.EQ(id),
	).One(context.Background(), r.db)
}

func (r *SqlDbChangeRequestRepo) GetAllChangeRequests() ([]*models.ChangeRequest, error) {
	crs, err := models.ChangeRequests().All(context.Background(), r.db)
	if err != nil {
		return nil, err
	}
	return crs, nil
}

func (r *SqlDbChangeRequestRepo) UpdateChangeRequest(cr *models.ChangeRequest) error {
	_, err := cr.Update(context.Background(), r.db, boil.Infer())
	return err
}

func (r *SqlDbChangeRequestRepo) AddTMTJob(cr *models.ChangeRequest, job *models.TMTJob) error {
	return cr.AddTMTJobs(context.Background(), r.db, true, job)
}

func (r *SqlDbChangeRequestRepo) AddRebaseJob(cr *models.ChangeRequest, job *models.RebaseJob) error {
	return cr.AddRebaseJobs(context.Background(), r.db, true, job)
}

func (r *SqlDbChangeRequestRepo) UpdateTMTJob(job *models.TMTJob) error {
	_, err := job.Update(context.Background(), r.db, boil.Infer())
	return err
}

func (r *SqlDbChangeRequestRepo) UpdateRebaseJob(job *models.RebaseJob) error {
	_, err := job.Update(context.Background(), r.db, boil.Infer())
	return err
}

func (r *SqlDbChangeRequestRepo) GetTMTJob(id int) (*models.TMTJob, error) {
	job, err := models.FindTMTJob(context.Background(), r.db, id)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (r *SqlDbChangeRequestRepo) GetRebaseJob(id int) (*models.RebaseJob, error) {
	job, err := models.FindRebaseJob(context.Background(), r.db, id)
	if err != nil {
		return nil, err
	}
	return job, nil
}
