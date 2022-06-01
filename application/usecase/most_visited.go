package usecase

import (
	"github.com/vinicch/shortener-go/application/port"
	"github.com/vinicch/shortener-go/domain"
)

// Gets the 10 most visited URLs
func MostVisited(getMostVisited port.GetMostVisited) ([]domain.Url, error) {
	return getMostVisited()
}
