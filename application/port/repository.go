package port

import "github.com/vinicch/shortener-go/domain"

// Gets information about an alias from the repository
type GetAlias func(string) (*domain.Url, error)

// Creates an alias record containing information about it and its URL
type CreateAlias func(*domain.Url) error

// Checks if a record for the given alias already exists
type DoesAliasExist func(string) bool
