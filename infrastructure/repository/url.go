package repository

import (
	"github.com/vinicch/shortener-go/application/port"
	"github.com/vinicch/shortener-go/domain"
	"gorm.io/gorm"
)

func GetAlias(db *gorm.DB) port.GetAlias {
	return func(alias string) (*domain.Url, error) {
		url := new(domain.Url)
		err := db.Where("alias = ?", alias).First(&url).Error

		return url, err
	}
}

func CreateAlias(db *gorm.DB) port.CreateAlias {
	return func(entity *domain.Url) error {
		return db.Create(entity).Error
	}
}

func DoesAliasExist(db *gorm.DB) port.DoesAliasExist {
	return func(alias string) bool {
		exists := false
		err := db.Where("alias = ?", alias).Find(&exists).Error
		if err != nil {
			return false
		}

		return exists
	}
}
