package main

import (
	"context"
	"database/sql"
	"embed"
	"flag"
	"fmt"
	"os"
	"sqlboiler-sb/models"
	"sqlboiler-sb/pkg/crspb"

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

	// Parse command line flags
	migrateFlag := flag.Bool("migrate", false, "run migrations")
	flag.Parse()

	if *migrateFlag {
		migrations := migrate.EmbedFileSystemMigrationSource{
			FileSystem: migrations,
			Root:       "migrate",
		}
		n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
		dieIf(err)
		println("Applied", n, "migrations")
	}
	flag.Parse()

	cr := &models.ChangeRequest{
		CreatedBy: "tyler@company.co", // get this from request
		Type:      crspb.ChangeRequestType_CRT_TMT.String(),
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
		CompanyDomain:           "some domain",
		UserEmail:               "alice",
		Status:                  crspb.JobStatus_PENDING.String(),
	}
	fmt.Println("before add:", tmtJob.ID)
	err = cr.AddTMTJobs(context.Background(), db, true, tmtJob)
	if err != nil {
		panic(err) // return error from gRPC handler
	}
	fmt.Println("after add:", tmtJob.ID)

	crs, _ := models.ChangeRequests(
		qm.Load(models.ChangeRequestRels.TMTJobs),
	).All(context.Background(), db)
	for _, cr := range crs {
		fmt.Println("change request:", cr.ID, cr.CreatedBy, cr.CreatedAt)
		for _, tj := range cr.R.TMTJobs {
			fmt.Println("  tmt job:", tj.StatusMessage, tj.CreatedAt, tj.CompletedAt)
		}
	}
}
