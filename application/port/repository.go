package port

import "github.com/vinicch/shortener-go/domain"

type DoesAliasExist func(string) bool
type CreateAlias func(domain.UrlAlias) error
type GetAlias func(string) (domain.UrlAlias, error)
