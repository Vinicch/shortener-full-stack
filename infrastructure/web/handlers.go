package web

import (
	"github.com/gin-gonic/gin"
	"github.com/vinicch/shortener-go/application/port"
)

func Create(createAlias port.CreateAlias, doesAliasExist port.DoesAliasExist) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func Retrieve(getAlias port.GetAlias) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func MostVisited() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
