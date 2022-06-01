package usecase

import (
	"errors"

	"github.com/vinicch/shortener-go/application/port"
	"github.com/vinicch/shortener-go/domain"
)

func Retrieve(getAlias port.GetURL, alias string) (string, error) {
	url, err := getAlias(alias)
	if err != nil {
		return "", err
	} else if url == nil {
		return "", errors.New(domain.ShortenedURLNotFound)
	}

	return url.Original, nil
}
