package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvCloudName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load env file")
	}
	cloudinaryName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	return cloudinaryName
}
func EnvCloudApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load env file")
	}
	cloudinaryApiKey := os.Getenv("CLOUDINARY_API_KEY")
	return cloudinaryApiKey
}

func EnvCloudApiSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load env file")
	}
	cloudinaryApiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	return cloudinaryApiSecret
}
func EnvCloudUploadFolder() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load env file")
	}
	cloudinaryUploadFolder := os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
	return cloudinaryUploadFolder
}
