package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var Log *logrus.Logger

func Init() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(os.Stdout)

	// Get log level from env (default to Info)
	level := os.Getenv("LOG_LEVEL")
	parsedLevel, err := logrus.ParseLevel(strings.ToLower(level))
	if err != nil {
		parsedLevel = logrus.InfoLevel
	}
	Log.SetLevel(parsedLevel)
	Log.Info("Log level set to:", parsedLevel)
}
