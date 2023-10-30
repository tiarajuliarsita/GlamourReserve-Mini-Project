package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	APPPORT int
	SECRET_KEY string
}

func LoadAPP() *AppConfig {
	var res = new(AppConfig)

	godotenv.Load(".env")

	if val, found := os.LookupEnv("APPPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid port value,", err.Error())
			return nil
		}
		// isRead = true
		res.APPPORT = port
	}
	
	if val, found := os.LookupEnv("SECRETKEY"); found {
		// isRead = true
		res.SECRET_KEY = val
	}
	return res
}
