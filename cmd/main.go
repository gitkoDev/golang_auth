package main

import (
	"github.com/gitkoDev/medods_task/cmd/api"
	"github.com/gitkoDev/medods_task/db"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Config stage
	if err := loadEnv(); err != nil {
		logrus.Fatal("error loading env file", err)
	}

	// DB connection stage
	db, err := db.PostgresConnection()
	if err != nil {
		logrus.Fatal("error connecting to db", err)
	}

	// Server connection stage
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		logrus.Fatal(err)
	}

}

func loadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	return nil
}
