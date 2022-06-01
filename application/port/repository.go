package port

import "github.com/vinicch/shortener-go/domain"

type GetAlias func(string) (*domain.Url, error)
type CreateAlias func(*domain.Url) error
type DoesAliasExist func(string) bool
