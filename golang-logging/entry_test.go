package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewEntry(t *testing.T) {
	logger := logrus.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.WithField("username", "janedoe").Info("Hello janedoe")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "johndoe")
	logrus.Info("Hello entry")
}
