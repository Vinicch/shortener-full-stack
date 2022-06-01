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
func Create(createURL port.CreateURL, doesAliasExist port.DoesAliasExist) gin.HandlerFunc {
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
		result, err := usecase.Shorten(createURL, doesAliasExist, address, alias)
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
func Retrieve(getURL port.GetURL, updateURL port.UpdateURL) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		alias := ctx.Param("alias")

		url, err := usecase.Retrieve(getURL, updateURL, alias)
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

// Gets the 10 most visited URLs
func MostVisited(getMostVisited port.GetMostVisited) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urls, err := usecase.MostVisited(getMostVisited)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		response := []mostVisitedResponse{}

		for _, url := range urls {
			response = append(response, mostVisitedResponse{
				Url:      url.Original,
				ShortUrl: url.Shortened,
				Visits:   url.Visits,
			})
		}

		ctx.JSON(http.StatusOK, response)
	}
}
