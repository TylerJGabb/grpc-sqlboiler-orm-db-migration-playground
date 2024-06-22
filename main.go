package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sqlboiler-sb/models"

	_ "github.com/lib/pq"
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

	users, err := models.Users().All(context.Background(), db)
	dieIf(err)
	for _, user  := range(users) {
		fmt.Println(user)
	}
	
}
