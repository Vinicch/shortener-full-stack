package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vinicch/shortener-go/infrastructure/logging"
	"github.com/vinicch/shortener-go/infrastructure/repository"
	"github.com/vinicch/shortener-go/infrastructure/web"
)

func main() {
	logging.Setup()

	db := repository.Connect()
	doesAliasExist := repository.DoesAliasExist(db)
	createURL := repository.CreateURL(db)
	getURL := repository.GetURL(db)
	updateURL := repository.UpdateURL(db)
	getMostVisited := repository.GetMostVisited(db)

	router := gin.Default()

	router.POST("/create", web.Create(createURL, doesAliasExist))
	router.GET("/:alias", web.Retrieve(getURL, updateURL))
	router.GET("/most-visited", web.MostVisited(getMostVisited))

	address := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	router.Run(address)
}
