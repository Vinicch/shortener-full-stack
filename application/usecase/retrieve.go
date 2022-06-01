package usecase

import (
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/vinicch/shortener-go/application/port"
	"github.com/vinicch/shortener-go/domain"
)

// Gets an URL that corresponds to a given alias
func Retrieve(getURL port.GetURL, updateURL port.UpdateURL, alias string) (string, error) {
	url, err := getURL(alias)
	if err != nil {
		log.Error().Err(err).Msg("Error getting URL info")
		return "", err
	} else if url == nil {
		return "", errors.New(domain.ShortenedURLNotFound)
	}

	url.Visits++

	err = updateURL(url)
	if err != nil {
		log.Error().Err(err).Msg("Error updating URL info")
		return "", err
	}

	return url.Original, nil
}
