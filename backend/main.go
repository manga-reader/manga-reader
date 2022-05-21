package main

import (
	"github.com/manga-reader/manga-reader/backend/router"
	"github.com/sirupsen/logrus"
)

func main() {
	r := router.SetupRouter(
		&router.Params{},
		&router.Options{},
	)
	logrus.Info("START LISTENING...")
	r.Run("localhost:6699")
}
