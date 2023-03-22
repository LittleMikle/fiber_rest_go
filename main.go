package main

import (
	"github.com/LittleMikle/fiber_rest_go/config"
	"github.com/LittleMikle/fiber_rest_go/pkg/handlers"
	"github.com/LittleMikle/fiber_rest_go/pkg/models"
	"github.com/LittleMikle/fiber_rest_go/storage"
	"github.com/gofiber/fiber/v2"

	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := config.CreateConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	db, err := storage.NewConnection(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	//DB = pointer on gorm.DB
	//db = value from our DataBase
	r := handlers.Repository{
		DB: db,
	}
	app := fiber.New()
	//access to struct inside method
	r.SetupRoutes(app)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
