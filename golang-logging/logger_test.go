package main

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()
	logger.Info("test logger")
	fmt.Println("Print Logger:", logger)
}

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("test trace")
	logger.Debug("test debug")
	logger.Info("test info")
	logger.Warn("test warn")
	logger.Error("test error")
	// logger.Fatal("test fatal")
	// logger.Panic("test panic")
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("test trace")
	logger.Debug("test debug")
	logger.Info("test info")
	logger.Warn("test warn")
	logger.Error("test error")
	// logger.Fatal("test fatal")
	// logger.Panic("test panic")
}
