package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	Port                string
	DatabaseUrl         string
	AccessTokenSecret   string
	CloudinaryCloudName string
	CloudinaryApiKey    string
	CloudinaryApiSecret string
}

func LoadEnvVariables() *EnvVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DATABASE_URL")
	accessToken := os.Getenv("ACCESS_TOKEN_SECRET")
	cloudname := os.Getenv("CLOUDINARYCLOUDNAME")
	cloudapikey := os.Getenv("CLOUDINARYAPIKEY")
	cloudapisecret := os.Getenv("CLOUDINARYAPISECRET")
	if port == "" || dbUrl == "" || accessToken == "" || cloudname == "" || cloudapikey == "" || cloudapisecret == "" {
		log.Fatal("Env variables are not all given")
	}

	return &EnvVariables{
		Port:                port,
		DatabaseUrl:         dbUrl,
		AccessTokenSecret:   accessToken,
		CloudinaryCloudName: cloudname,
		CloudinaryApiKey:    cloudapikey,
		CloudinaryApiSecret: cloudapisecret,
	}
}
