package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CloudBucket() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load env file")
	}
	bucketName := os.Getenv("BUCKET_NAME")
	return bucketName
}

func CloudAccount() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load env file")
	}
	accountId := os.Getenv("ACCOUNT_ID")
	return accountId
}

func CloudKeySecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load env file")
	}
	accesKeyId := os.Getenv("ACCESS_KEY_SECRET")
	return accesKeyId
}
func CloudKeyId() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load env file")
	}
	cloudKeyId := os.Getenv("ACCESS_KEY_ID")
	return cloudKeyId
}
