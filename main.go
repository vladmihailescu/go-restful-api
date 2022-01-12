package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vladmihailescu/go-restful-api/database"
	"github.com/vladmihailescu/go-restful-api/routes"
)

var (
	PORT = ":8080"
)

func main() {
	log.SetLevel(log.DebugLevel)

	if err := database.InitDatabase(); err != nil {
		log.Fatalf("unable to init database: %v", err)
	}

	router := gin.New()

	routes.InitUserRoutes(router)

	log.Fatal(router.Run(PORT))
}
