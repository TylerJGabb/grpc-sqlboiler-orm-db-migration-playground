package grpcserver

import (
	"context"
	"database/sql"
	"sqlboiler-sb/jobs"
	"sqlboiler-sb/models"
	"sqlboiler-sb/pkg/crspb"
	"sqlboiler-sb/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func requestTypeHelper(crType string) crspb.ChangeRequestType {
	typeInt, ok := crspb.ChangeRequestType_value[crType]
	if !ok {
		return crspb.ChangeRequestType_UNKNOWN
	}
	return crspb.ChangeRequestType(typeInt)
}

type Server struct {
	repo   repository.ChangeRequestRepository
	runner jobs.CloudRunJobRunner
	crspb.UnimplementedChangeRequestServiceServer
}

func (s *Server) ReportDefaultBranchUpdated(
	ctx context.Context,
	req *crspb.ReportDefaultBranchUpdatedRequest,
) (*crspb.ReportDefaultBranchUpdatedResponse, error) {
	crs, err := s.repo.GetAllChangeRequests()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	for _, cr := range crs {
		rebaseJobModel := &models.RebaseJob{
			Status: crspb.JobStatus_PENDING.String(),
		}
		err := s.repo.AddRebaseJob(cr, rebaseJobModel)
		if err != nil {
			// continue, because we want to update all CRs
			// but log the error
		}
		job := &jobs.RebaseJob{
			JobId:           rebaseJobModel.ID,
			ChangeRequestId: cr.ID,
		}
		_, err = s.runner.RunJob(job)
		if err != nil {
			// continue, because we want to update all CRs
			// but log the error
		}
	}
	return &crspb.ReportDefaultBranchUpdatedResponse{
		Success: true,
	}, nil
}

func NewServer(db *sql.DB) *Server {
	return &Server{}
}

func (s *Server) CreateTMTProject(
	ctx context.Context,
	req *crspb.CreateTMTProjectRequest,
) (*crspb.CreateTMTProjectResponse, error) {

	cr := &models.ChangeRequest{
		CreatedBy: req.UserEmail,
		Type:      crspb.ChangeRequestType_TMT.String(),
	}
	err := s.repo.CreateChangeRequest(cr)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	tmtJob := &models.TMTJob{
		ProjectName:             req.ProjectName,
		OrchestrationRepository: req.OrchestrationRepository,
		Application:             req.Application,
		DV01Domain:              req.Dv01Domain,
		UserEmail:               req.UserEmail,
		Status:                  crspb.JobStatus_PENDING.String(),
	}

	err = s.repo.AddTMTJob(cr, tmtJob)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	job := &jobs.CreateTMTPR{
		JobId:           tmtJob.ID,
		ChangeRequestId: cr.ID,
	}

	_, err = s.runner.RunJob(job)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &crspb.CreateTMTProjectResponse{
		Success:         true,
		ChangeRequestId: int32(cr.ID),
	}, nil
}

func (s *Server) GetChangeRequest(
	ctx context.Context,
	req *crspb.GetChangeRequestRequest,
) (*crspb.ChangeRequest, error) {
	cr, err := s.repo.GetChangeRequest(int(req.ChangeRequestId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &crspb.ChangeRequest{
		Id:             int32(cr.ID),
		BranchName:     "cr.BranchName",
		PullRequestUrl: cr.GithubPRURL.String,
		PullRequestId:  cr.GithubPRID.String,
		CreatedBy:      cr.CreatedBy,
		Type:           requestTypeHelper(cr.Type),
	}, nil
}
