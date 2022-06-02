package repository

import (
	"errors"

	"github.com/vinicch/shortener-go/internal/application/port"
	"github.com/vinicch/shortener-go/internal/domain"
	"gorm.io/gorm"
)

func getURL(db *gorm.DB) port.GetURL {
	return func(alias string) (*domain.Url, error) {
		url := new(domain.Url)
		err := db.Where("alias = ?", alias).First(&url).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return url, err
	}
}

func getMostVisited(db *gorm.DB) port.GetMostVisited {
	return func() ([]domain.Url, error) {
		urls := make([]domain.Url, 10)
		err := db.Order("visits DESC").Limit(10).Find(&urls).Error

		return urls, err
	}
}

func createURL(db *gorm.DB) port.CreateURL {
	return func(entity *domain.Url) error {
		return db.Create(entity).Error
	}
}

func updateURL(db *gorm.DB) port.UpdateURL {
	return func(url *domain.Url) error {
		return db.Save(&url).Error
	}
}

func doesAliasExist(db *gorm.DB) port.DoesAliasExist {
	return func(alias string) bool {
		var count int64
		err := db.Model(&domain.Url{}).Where("alias = ?", alias).Count(&count).Error
		if err != nil {
			return false
		}

		return count > 0
	}
}
