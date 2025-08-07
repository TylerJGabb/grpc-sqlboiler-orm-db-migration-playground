package jobs

import (
	run "cloud.google.com/go/run/apiv2"
	"cloud.google.com/go/run/apiv2/runpb"
)

type Job interface {
	ToRunJobRequest(jobFqn string) *runpb.RunJobRequest
}

type CloudRunJobRunner interface {
	RunJob(job Job) (*run.RunJobOperation, error)
}
