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
	createAlias := repository.CreateAlias(db)
	getAlias := repository.GetAlias(db)

	router := gin.Default()

	router.POST("/create", web.Create(createAlias, doesAliasExist))
	router.GET("/:alias", web.Retrieve(getAlias))
	router.GET("/mostVisited", web.MostVisited())

	address := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	router.Run(address)
}
