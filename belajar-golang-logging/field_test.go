package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("email", "nurfirdaus@gmail.com").Info("test")
	logger.WithField("email", "nurfirdaus@gmail.com").Error("test error")
}
