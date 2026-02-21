package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	config := LoadConfig()

	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("Starting application")

	router := SetupRouter(config.Port)

	addr := fmt.Sprintf(":%d", config.Port)
	router.Run(addr)
}