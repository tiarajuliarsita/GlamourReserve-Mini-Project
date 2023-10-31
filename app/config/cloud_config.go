package config

import (
	"os"

	"github.com/joho/godotenv"
)

func CloudBucket() string {
	godotenv.Load()
	// if err != nil {
	// 	log.Fatal("can't load env file")
	// }
	bucketName := os.Getenv("BUCKET_NAME")
	return bucketName
}

func CloudAccount() string {
	godotenv.Load()
	// if err != nil {
	// 	log.Fatal("can't load env file")
	// }
	accountId := os.Getenv("ACCOUNT_ID")
	return accountId
}

func CloudKeySecret() string {
	godotenv.Load()

	accesKeyId := os.Getenv("ACCESS_KEY_SECRET")
	return accesKeyId
}
func CloudKeyId() string {
	godotenv.Load()

	cloudKeyId := os.Getenv("ACCESS_KEY_ID")
	return cloudKeyId
}
