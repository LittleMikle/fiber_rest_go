package config

import (
	"github.com/LittleMikle/fiber_rest_go/storage"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

func CreateConfig() (Conf *storage.Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return nil, err
	} else {
		log.Info().Msg("Config created successfully")
	}

	Conf = &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	return Conf, nil
}
