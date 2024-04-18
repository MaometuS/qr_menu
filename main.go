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
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := router.NewRouter(controller.NewController(config, newDB(config.DbUrl)))

	fmt.Println("Serving on 4000")
	http.ListenAndServe(":4000", app)
}

func getConfig() (*entity.Config, error) {
	mailPort, err := strconv.ParseInt(os.Getenv("mail_port"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &entity.Config{
		DbUrl:         os.Getenv("db_url"),
		UploadFolder:  os.Getenv("upload_folder"),
		JWTSignString: os.Getenv("jwt_sign_string"),
		MailFrom:      os.Getenv("mail_from"),
		MailPassword:  os.Getenv("mail_password"),
		MailAddress:   os.Getenv("mail_address"),
		MailPort:      mailPort,
	}, nil
}

func newDB(url string) entity.PgxIface {
	db, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
