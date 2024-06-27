package jobs

import "cloud.google.com/go/run/apiv2/runpb"

type CreateTMTPR struct {
	ProjectName             string
	OrchestrationRepository string
	Application             string
	Dv01Domain              string
	UserEmail               string
}

// see example at https://github.com/dv01-inc/dynamic-sandbox-api/blob/main/messaging/cloudrun/cloudrun.go#L47
func (c *CreateTMTPR) ToRunJobRequest(jobFqn string) *runpb.RunJobRequest {
	//...
	// build run job request with appropriate command line argument overrides
	return nil
}
