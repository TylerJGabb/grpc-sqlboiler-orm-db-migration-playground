package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"os"
	"sqlboiler-sb/models"

	_ "github.com/lib/pq"
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

	customers, _ := models.Customers(qm.Load("Invoices.Products")).All(context.Background(), db)
	for _, c := range customers {
		fmt.Println(c.Name)
		for _, i := range c.R.Invoices {
			fmt.Println("  invoice:", i.ID)
			for _, p := range i.R.Products {
				fmt.Println("    product:", p.Sku)
			}
		}
	}
}
