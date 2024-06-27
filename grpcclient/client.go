package grpcclient

import (
	"sqlboiler-sb/pkg/crspb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewClient(serverAddr string) (crspb.ChangeRequestServiceClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
	}
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		return nil, err
	}
	client := crspb.NewChangeRequestServiceClient(conn)
	return client, nil
}
