package main

import (
	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/router"
	"github.com/sirupsen/logrus"
)

func main() {
	db := database.Connect()

	r := router.SetupRouter(
		&router.Params{
			Database: db,
		},
		&router.Options{},
	)
	logrus.Info("START LISTENING...")
	r.Run("localhost:6699")
}
