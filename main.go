package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"os"
	"sqlboiler-sb/models"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrate/*
var migrations embed.FS

//go:generate sqlboiler psql --wipe

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	dieIf(os.Setenv("PGPASSWORD", "pass"))

	db, err := sql.Open("postgres", "dbname=db user=user sslmode=disable")
	dieIf(err)

	dieIf(db.Ping())

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrations,
		Root:       "migrate",
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	dieIf(err)
	println("Applied", n, "migrations")

	cr := &models.ChangeRequest{
		CreatedBy: "tyler@dv01.co", // get this from request
		Type:      models.ChangeRequestTypeTMTProject,
	}
	err = cr.Insert(
		context.Background(),
		db,
		boil.Infer(),
	)
	if err != nil {
		panic(err) // return error from gRPC handler
	}
	tmtJob := &models.TMTJob{
		ProjectName:             "foobarbaz",
		OrchestrationRepository: "some repo",
		Application:             "some app",
		DV01Domain:              "some domain",
		UserEmail:               "alice",
		Status:                  models.JobStatusPending,
	}
	err = cr.AddTMTJobs(context.Background(), db, true, tmtJob)
	if err != nil {
		panic(err) // return error from gRPC handler
	}

	crs, _ := models.ChangeRequests(
		qm.Load(models.ChangeRequestRels.TMTJobs),
	).All(context.Background(), db)
	for _, cr := range crs {
		fmt.Println("change request:", cr.ID, cr.CreatedBy, cr.CreatedAt)
		if cr.Type == models.ChangeRequestTypeTMTProject {
			for _, tj := range cr.R.TMTJobs {
				fmt.Println("  tmt job:", tj.StatusMessage, tj.CreatedAt, tj.CompletedAt)
			}
		}
	}
}
