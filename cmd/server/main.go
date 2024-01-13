package main

import (
	"simple-to-do/internal/app"
	"simple-to-do/internal/config"
	"simple-to-do/internal/utils/constants"
	"simple-to-do/pkg/logger"

	"github.com/sirupsen/logrus"
)

func init() {
	if err := config.Read(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
	logger.Info("configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})

	if config.IsAppProd() {
		logger.Info("**Production Environment**", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
}

func main() {
	s, err := app.InitializeApp()
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	}
	logger.Info("initialize app", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

	err = s.ListenAndServe()
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	}
}
