package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sqlboiler-sb/models"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

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
	
	customer1, err := models.Customers(qm.Where("id = ?", 1)).One(context.Background(), db)
	dieIf(err)

	fmt.Println("customer1:", customer1.Name)

	invoice, err := customer1.Invoices().One(context.Background(), db)
	dieIf(err)

	fmt.Println("  invoice:", invoice.ID)

	invoiceItems, err := invoice.Products().All(context.Background(), db)
	dieIf(err)

	for _, item := range invoiceItems {
		fmt.Println("    item:", item.Sku, item)
	}

	customer2, err := models.Customers(qm.Where("id = ?", 2)).One(context.Background(), db)
	dieIf(err)

	fmt.Println("customer2:", customer2.Name)

	invoice, err = customer2.Invoices().One(context.Background(), db)
	dieIf(err)

	fmt.Println("  invoice:", invoice.ID)

	invoiceItems, err = invoice.Products().All(context.Background(), db)
	dieIf(err)

	for _, item := range invoiceItems {
		fmt.Println("    item:", item.Sku, item)
	}
}
