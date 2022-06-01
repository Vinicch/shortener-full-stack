package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vinicch/shortener-go/infrastructure/logging"
	"github.com/vinicch/shortener-go/infrastructure/repository"
)

func main() {
	logging.Setup()

	db := repository.Connect()
	_ = repository.DoesAliasExist(db)
	_ = repository.CreateAlias(db)
	_ = repository.GetAlias(db)

	router := gin.Default()

	router.POST("/create")
	router.GET("/{alias}")
	router.GET("/mostVisited")

	address := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	router.Run(address)
}
