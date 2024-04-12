package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/interface/controller"
	"gitlab.com/maometusu/qr_menu/internal/router"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := router.NewRouter(controller.NewController(newDB()))

	fmt.Println("Serving on 4000")
	http.ListenAndServe(":4000", app)
}

func newDB() entity.PgxIface {
	db, err := pgxpool.New(context.Background(), os.Getenv("db_url"))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
