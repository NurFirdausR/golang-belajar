package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logrus := logrus.New()

	logrus.Println("hello world")
	logrus.Trace("This is trace")
	logrus.Debug("This is Debug")
	logrus.Info("This is info")
	logrus.Warn("This is warning")
	logrus.Error("This is Error")
}
