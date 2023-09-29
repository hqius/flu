package config

import "github.com/sirupsen/logrus"

func initLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
