package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/manga-reader/manga-reader/backend/config"
	"github.com/manga-reader/manga-reader/backend/database"
	"github.com/manga-reader/manga-reader/backend/router"
	"github.com/manga-reader/manga-reader/backend/usecases"
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
	err := config.LoadConfiguration("./config.json")
	if err != nil {
		logrus.Fatal(err)
	}

	db := database.NewDatabase(
		database.Default_Host,
		database.Default_Port,
		database.Default_User,
		database.Default_Password,
		database.Default_Dbname,
	)
	if err = db.Connect(); err != nil {
		logrus.Fatal("failed to init db")
	}
	u := usecases.NewUsecase(db)

	r := router.SetupRouter(
		&router.Params{
			Usecase: u,
		},
	)
	logrus.Info("START LISTENING...")
	r.Run(fmt.Sprintf("%s:%s", config.Cfg.Connection.ExportHost, config.Cfg.Connection.ExportPort))
}
