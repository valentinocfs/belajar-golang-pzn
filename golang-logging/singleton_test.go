package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Logrus Info")
	logrus.Warn("Logrus Warn")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Logrus Info")
	logrus.Warn("Logrus Warn")
}
