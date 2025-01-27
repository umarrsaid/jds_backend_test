package main

import (
	"be-golang/config"
	"be-golang/db"
	"be-golang/router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.NewDB(cfg.DBURI)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	router.InitRouters(app, dbConn, cfg)

	// Load port from environment variable, default to ":3001" if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	// Start the server
	log.Printf("Server is running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
