package initializers

import "github.com/joho/godotenv"

func LoadEnvVariable() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}
