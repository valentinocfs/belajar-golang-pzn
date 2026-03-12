package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "johndoe").Info("Hello John Doe")
	logger.WithField("username", "janedoe").
		WithField("email", "janedoe@gmail.com").
		Info("Hello Jane Doe")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "johndoe",
		"email":    "johndoe@gmail.com",
		"theme":    "dark",
	}).Info("Hello John Doe")
}
