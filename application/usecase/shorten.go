package usecase

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vinicch/shortener-go/application/port"
	"github.com/vinicch/shortener-go/domain"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Shorten(createAlias port.CreateAlias, doesAliasExist port.DoesAliasExist,
	url, alias string) (domain.Url, error) {

	if strings.TrimSpace(alias) != "" {
		if doesAliasExist(alias) {
			return domain.Url{}, errors.New("CUSTOM ALIAS ALREADY EXISTS")
		}
	} else {
		alias = generateAlias(doesAliasExist)
	}

	host := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	entity := domain.Url{
		Id:        uuid.NewString(),
		Alias:     alias,
		Original:  url,
		Shortened: fmt.Sprintf("%s/%s", host, alias),
	}

	err := createAlias(&entity)
	if err != nil {
		return domain.Url{}, err
	}

	return entity, nil
}

func generateAlias(doesAliasExist port.DoesAliasExist) string {
	rand.Seed(time.Now().UnixNano())

	alias := make([]byte, 6)
	for i := range alias {
		index := rand.Intn(len(chars))
		alias[i] = chars[index]
	}

	if doesAliasExist(string(alias)) {
		return generateAlias(doesAliasExist)
	}

	return string(alias)
}