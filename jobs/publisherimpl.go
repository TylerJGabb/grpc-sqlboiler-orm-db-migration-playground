package jobs

import (
	"context"

	run "cloud.google.com/go/run/apiv2"
	"cloud.google.com/go/run/apiv2/runpb"
	"github.com/googleapis/gax-go/v2"
)

// https://github.com/googleapis/google-cloud-go/blob/main/testing.md
// https://pkg.go.dev/cloud.google.com/go/run/apiv2

// Create our own interface, which is satiasfied by the *run.JobsClient in cloud.google.com/go/run/apiv2
// to enable mocking in tests, since the *run.JobsClient does not implement an interface
type CloudRunJobsClient interface {
	RunJob(ctx context.Context, in *runpb.RunJobRequest, opts ...gax.CallOption) (*run.RunJobOperation, error)
}

type CloudRunJobPublisherImpl struct {
	jobFqn string
	client CloudRunJobsClient
	ctx    context.Context
}

func (c *CloudRunJobPublisherImpl) Publish(job Job) (*run.RunJobOperation, error) {
	req := job.ToRunJobRequest()
	op, err := c.client.RunJob(c.ctx, req)
	return op, err
}
