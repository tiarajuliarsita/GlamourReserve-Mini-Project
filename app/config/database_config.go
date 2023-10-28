package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type DBConfig struct {
	DBPORT      int
	DBHOST      string
	DBUSER      string
	DBPASSWORD string
	DBNAME      string
}



func LoadDB() *DBConfig {
	var res = new(DBConfig)

	godotenv.Load(".env")
	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid db port value,", err.Error())
			return nil
		}
		// isRead = true
		res.DBPORT = port
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		// isRead = true
		res.DBHOST = val
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		// isRead = true
		res.DBUSER = val
	}

	if val, found := os.LookupEnv("DBPASSWORD"); found {
		// isRead = true
		res.DBPASSWORD = val
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		// isRead = true
		res.DBNAME = val
	}

	return res

}
