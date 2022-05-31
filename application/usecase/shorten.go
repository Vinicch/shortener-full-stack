package usecase

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vinicch/shortener-go/application/port"
	"github.com/vinicch/shortener-go/domain"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Shorten(createAlias port.CreateAlias, doesAliasExist port.DoesAliasExist,
	url, alias string) (domain.UrlAlias, error) {

	if strings.TrimSpace(alias) != "" {
		if doesAliasExist(alias) {
			return domain.UrlAlias{}, errors.New("CUSTOM ALIAS ALREADY EXISTS")
		}
	} else {
		alias = generateAlias(doesAliasExist)
	}

	entity := domain.UrlAlias{
		Id:    uuid.NewString(),
		Alias: alias,
		Url:   url,
	}

	err := createAlias(entity)
	if err != nil {
		return domain.UrlAlias{}, err
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
