package jobs

import (
	"strconv"

	"cloud.google.com/go/run/apiv2/runpb"
)

type CreateTMTPR struct {
	JobId           int
	ChangeRequestId int
}

// see example at https://github.com/company-inc/dynamic-sandbox-api/blob/main/messaging/cloudrun/cloudrun.go#L47
func (c *CreateTMTPR) ToRunJobRequest(jobFqn string) *runpb.RunJobRequest {
	return &runpb.RunJobRequest{
		Name: jobFqn,
		Overrides: &runpb.RunJobRequest_Overrides{
			ContainerOverrides: []*runpb.RunJobRequest_Overrides_ContainerOverride{
				{
					Env: []*runpb.EnvVar{
						{
							Name: "JOB_ID",
							Values: &runpb.EnvVar_Value{
								Value: strconv.Itoa(c.JobId),
							},
						},
						{
							Name: "CHANGE_REQUEST_ID",
							Values: &runpb.EnvVar_Value{
								Value: strconv.Itoa(c.ChangeRequestId),
							},
						},
						{
							Name: "TYPE",
							Values: &runpb.EnvVar_Value{
								// ideally, get this from the enum defined in the contract.
								// it's hard coded here for demonstration
								Value: "TMT",
							},
						},
					},
				},
			},
		},
	}
}
