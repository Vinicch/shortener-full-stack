package web

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vinicch/shortener-go/application/port"
	"github.com/vinicch/shortener-go/application/usecase"
	"github.com/vinicch/shortener-go/domain"
)

// Creates a shortened version of a provided URL
func Create(createAlias port.CreateAlias, doesAliasExist port.DoesAliasExist) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		address := ctx.Query("url")
		alias := ctx.Query("CUSTOM_ALIAS")

		// Validates URL
		if strings.TrimSpace(address) == "" {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("'url' not informed"))
			return
		}

		_, err := url.Parse(address)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid URL"))
			return
		}

		// Shortens the URL and measures execution time
		start := time.Now().UTC()
		result, err := usecase.Shorten(createAlias, doesAliasExist, address, alias)
		if err != nil {
			if err.Error() == domain.AliasAlreadyExists {
				ctx.JSON(http.StatusConflict, errorResponse{Code: "001", Description: err.Error()})
				return
			}

			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		elapsed := time.Since(start).String()

		// Returns the shortened URL along with related data
		response := createResponse{
			Alias:       result.Alias,
			Original:    result.Original,
			Shortened:   result.Shortened,
			ElapsedTime: elapsed,
		}

		ctx.JSON(http.StatusCreated, response)
	}
}

// Gets the URL that corresponds to the alias and redirects to it
func Retrieve(getAlias port.GetAlias) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		alias := ctx.Param("alias")

		url, err := usecase.Retrieve(getAlias, alias)
		if err != nil {
			if err.Error() == domain.ShortenedURLNotFound {
				ctx.JSON(http.StatusNotFound, errorResponse{Code: "002", Description: err.Error()})
				return
			}

			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, url)
	}
}

func MostVisited() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
