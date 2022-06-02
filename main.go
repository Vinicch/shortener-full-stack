package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vinicch/shortener-go/infrastructure/logging"
	"github.com/vinicch/shortener-go/infrastructure/repository"
	"github.com/vinicch/shortener-go/infrastructure/web"
)

func main() {
	logging.Setup()

	urlFunctions := repository.MakeURLFunctions()
	router := gin.Default()

	router.POST("/create", web.Create(urlFunctions.CreateURL, urlFunctions.DoesAliasExist))
	router.GET("/url/:alias", web.Retrieve(urlFunctions.GetURL, urlFunctions.UpdateURL))
	router.GET("/most-visited", web.MostVisited(urlFunctions.GetMostVisited))

	router.Run()
}
