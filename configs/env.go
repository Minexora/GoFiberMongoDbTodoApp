package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error .env")
	}
	mongoUri := os.Getenv("MONGO_URI")
	return mongoUri
}
