package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/router"
	"github.com/sirupsen/logrus"
)

const logFile = false

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)

	if logFile {
		fileName := fmt.Sprintf("manga-reader-%s.log", time.Now().Format(time.RFC3339))
		logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("fail to create log file")
		}
		mw := io.MultiWriter(os.Stdout, logFile)
		logrus.SetOutput(mw)
	}
}

func main() {
	db := database.Connect()

	r := router.SetupRouter(
		&router.Params{
			Database: db,
		},
	)
	logrus.Info("START LISTENING...")
	r.Run("localhost:6699")
}
