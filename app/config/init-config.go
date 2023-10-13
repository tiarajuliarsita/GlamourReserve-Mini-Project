package config

import "github.com/sirupsen/logrus"

func InitConfig() (*AppConfig, *DBConfig) {

	db := LoadDB()
	app := LoadAPP()

	if db == nil || app == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil, nil
	}

	return app, db
}
