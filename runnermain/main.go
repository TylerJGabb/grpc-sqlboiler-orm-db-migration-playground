package runnermain

import (
	"context"
	"sqlboiler-sb/envconfig"
	"sqlboiler-sb/grpcclient"
	"sqlboiler-sb/pkg/crspb"
)

func main() {
	grpcConfig, err := envconfig.LoadGrpcClientConfig()
	if err != nil {
		panic(err) // don't actually panic, log useful messages
	}

	client, err := grpcclient.NewClient(grpcConfig.GrpcServerAddress())
	if err != nil {
		panic(err) // don't actually panic, log useful messages
	}

	runConfig, err := envconfig.LoadRunConfig()
	if err != nil {
		panic(err) // don't actually panic, log useful messages
	}

	changeRequest, err := client.GetChangeRequest(context.Background(), &crspb.GetChangeRequestRequest{
		ChangeRequestId: int32(runConfig.ChangeRequestId),
	})

	if err != nil {
		panic(err) // don't actually panic, log useful messages
	}

	// I'll leave the implementation details up to you, but in general you'll want to:
	// git pull
	// git checkout changeRequest.BranchName
	// git rebase origin/main
	// git push --force
	err = performRebase(changeRequest)

	if err != nil {
		client.UpdateRebaseJobStatus(context.Background(), &crspb.UpdateJobStatusRequest{
			JobId:         int32(runConfig.JobId),
			Status:        crspb.JobStatus_FAILED,
			StatusMessage: err.Error(),
		})
	} else {
		client.UpdateRebaseJobStatus(context.Background(), &crspb.UpdateJobStatusRequest{
			JobId:  int32(runConfig.JobId),
			Status: crspb.JobStatus_COMPLETED,
		})
	}

	grpcclient.NewClient("localhost:8080")
}
