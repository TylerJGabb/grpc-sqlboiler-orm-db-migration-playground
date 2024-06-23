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
	
	users, err := models.Users().All(context.Background(), db)
	dieIf(err)

	for _, user  := range(users) {
		fmt.Println(user)
	}

	doesIt, err := models.Videos(qm.Where("id = ?", 1)).Exists(context.Background(), db)
	dieIf(err)

	fmt.Println("does it?", doesIt)

	user, err := models.Users().One(context.Background(), db)
	dieIf(err)

	fmt.Println("user: ", user.ID, user.Name)

	nVideos, err := user.Videos().Count(context.Background(), db)
	dieIf(err)

	fmt.Println("  nVideos:", nVideos)

	vid1, vid2 := &models.Video{Name: "a"}, &models.Video{Name: "b"}
	err = user.AddVideos(context.Background(), db, true, vid1, vid2)
	dieIf(err)

	fmt.Println("user:", user.Name)
	nVideos, err = user.Videos().Count(context.Background(), db)
	dieIf(err)

	fmt.Println("  nVideos:", nVideos)
}
