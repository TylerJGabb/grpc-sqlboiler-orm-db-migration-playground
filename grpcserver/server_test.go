package grpcserver_test

import (
	"context"
	"sqlboiler-sb/grpcserver"
	"sqlboiler-sb/pkg/crspb"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_Server(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()
	server := grpcserver.NewServer(db)
	res, err := server.GetChangeRequest(context.Background(), &crspb.GetChangeRequestRequest{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	t.Log(res)
	t.Log(mock)

}
