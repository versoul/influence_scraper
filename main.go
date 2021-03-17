package main

import (
	"influence_scraper/api"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const HOST = ":9090"

func main() {
	r := gin.Default()

	rateHandler := api.NewRateHandler()
	rateHandler.MountRoutes(r)

	r.Use(static.Serve("/", static.LocalFile("./static", true)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})

	log.WithFields(log.Fields{
		"addr": HOST,
	}).Info("Server started...")

	if err := r.Run(HOST); err != nil {
		log.Fatal(err)
	}
}
