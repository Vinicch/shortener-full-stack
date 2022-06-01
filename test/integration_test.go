package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/vinicch/shortener-go/infrastructure/repository"
	"github.com/vinicch/shortener-go/infrastructure/web"
)

func setup() *gin.Engine {
	godotenv.Load("../.env")

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

	return router
}

func TestCreate(t *testing.T) {
	router := setup()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/create?url=http://hostname.com/long/url/path", nil)

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestCreateConflict(t *testing.T) {
	router := setup()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/create?url=http://hostname.com/long/url/path&CUSTOM_ALIAS=test", nil)

	// Runs twice to guarantee conflict
	router.ServeHTTP(recorder, req)
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusConflict, recorder.Code)
}

func TestRetrieve(t *testing.T) {
	router := setup()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusMovedPermanently, recorder.Code)
}

func TestRetrieveNotFound(t *testing.T) {
	router := setup()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/qwertyuiop", nil)

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestMostVisited(t *testing.T) {
	router := setup()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/most-visited", nil)

	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
