package port

import "github.com/vinicch/shortener-go/domain"

// Gets information about an URL from the repository
type GetURL func(string) (*domain.Url, error)

// Creates an URL record containing information about it and its alias
type CreateURL func(*domain.Url) error

// Checks if a record for the given alias already exists
type DoesAliasExist func(string) bool
