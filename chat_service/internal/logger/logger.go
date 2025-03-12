package logger

import (
	"chat_service/internal/config"
	"os"

	"github.com/sirupsen/logrus"
)

func SetUpLogger(cfg *config.Config) (*logrus.Logger, *os.File) {
	log := logrus.New()
	var file *os.File = nil

	switch cfg.LogLevel {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	}

	switch cfg.LogOutput {
	case "stdout":
		logrus.SetOutput(os.Stdout)
		log.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	case "file":
		if cfg.LogFilePath == "" {
			logrus.Fatal("Log output type is file, but file path is not specified. Specify filepath.")
		}

		file, err := os.OpenFile(cfg.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		if err != nil {
			logrus.Fatal("Unable to open log file:", err)
		}

		logrus.SetOutput(file)
	}

	return log, file
}

func Close(file *os.File) {
	if file != nil {
		file.Close()
	}
}
