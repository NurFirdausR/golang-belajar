package belajar_golang_logging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	logger.SetOutput(file)

	logger.Info("hello logging info")
	logger.Warn("hello logging warn")
	logger.Error("hello logging err")
}
