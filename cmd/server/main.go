package main

import (
	"simple-to-do/internal/app"
	"simple-to-do/internal/config"
	"simple-to-do/internal/utils/constants"
	pkg_logger "simple-to-do/pkg/logger"

	"github.com/sirupsen/logrus"
)

func init() {
	if err := config.InitializeAppConfig(); err != nil {
		pkg_logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
	pkg_logger.Info("configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})

	if config.IsAppProd() {
		pkg_logger.Info("**Production Environment**", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
}

func main() {
	s, err := app.InitializeApp()
	if err != nil {
		pkg_logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	}
	pkg_logger.Info("initialize app", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

	err = s.ListenAndServe()
	if err != nil {
		pkg_logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	}
}
