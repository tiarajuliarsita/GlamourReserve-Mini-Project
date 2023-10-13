package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	APPPORT int
	SECRETKEY string
}

func LoadAPP() *AppConfig {
	var res = new(AppConfig)

	godotenv.Load(".env")

	// var isRead = false
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
		res.SECRETKEY = val
	}
	return res
}
