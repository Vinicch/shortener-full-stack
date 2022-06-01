package usecase

import (
	"github.com/vinicch/shortener-go/application/port"
	"github.com/vinicch/shortener-go/domain"
)

func Retrieve(getAlias port.GetAlias, alias string) (*domain.Url, error) {
	return getAlias(alias)
}
