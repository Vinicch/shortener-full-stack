package usecase

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/vinicch/shortener-go/application/port"
	"github.com/vinicch/shortener-go/domain"
)

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
	random := rand.Int63()
	alias := strconv.FormatInt(random, 36)
	alias = alias[0:5]

	if doesAliasExist(alias) {
		return generateAlias(doesAliasExist)
	}

	return alias
}
