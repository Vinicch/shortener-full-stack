package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinicch/shortener-go/domain"
)

func TestRetrieve(t *testing.T) {
	alias := "test"
	url := "http://hostname.com/long/url/path"

	getURLStub := func(string) (*domain.Url, error) {
		return &domain.Url{
			Alias:    alias,
			Original: url,
		}, nil
	}

	updateURLStub := func(*domain.Url) error { return nil }
	result, err := Retrieve(getURLStub, updateURLStub, alias)

	if assert.NoError(t, err) {
		assert.Equal(t, url, result)
	}
}

func TestRetrieveInput(t *testing.T) {
	alias := ""
	getURLStub := func(string) (*domain.Url, error) { return nil, nil }
	updateURLStub := func(*domain.Url) error { return nil }

	_, err := Retrieve(getURLStub, updateURLStub, alias)

	assert.EqualError(t, err, domain.AliasNotInformed)
}

func TestRetrieveDeps(t *testing.T) {
	alias := "test"
	getURLErr := "getURLError"
	updateURLErr := "updateURLError"

	getURLStub := func(string) (*domain.Url, error) { return &domain.Url{}, nil }
	updateURLStub := func(*domain.Url) error { return errors.New(updateURLErr) }
	_, err := Retrieve(getURLStub, updateURLStub, alias)
	assert.EqualError(t, err, updateURLErr)

	getURLStub = func(string) (*domain.Url, error) { return nil, errors.New(getURLErr) }
	_, err = Retrieve(getURLStub, updateURLStub, alias)
	assert.EqualError(t, err, getURLErr)

	getURLStub = func(string) (*domain.Url, error) { return nil, nil }
	_, err = Retrieve(getURLStub, updateURLStub, alias)
	assert.EqualError(t, err, domain.ShortenedURLNotFound)
}
