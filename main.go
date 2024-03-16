package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/maometusu/qr_menu/internal/controller"
	"gitlab.com/maometusu/qr_menu/internal/router"
	"log"
	"net/http"
	"os"
)

func main() {
	_, err := pgxpool.New(context.Background(), os.Getenv("db_url"))
	if err != nil {
		log.Fatal(err)
	}

	app := router.NewRouter(controller.NewController())

	fmt.Println("Serving on 4000")
	http.ListenAndServe(":4000", app)
}
