package dao

import (
	"gorm.io/gorm"
	"vizzy/entity"
)

type Page1Dao struct {
	DB *gorm.DB
}

func (d Page1Dao) Create(form entity.Page1) error {
	return d.DB.Model(&entity.Page1{}).Create(&form).Error
}

func (d Page1Dao) Query() (entity.Page1, error) {
	var record entity.Page1
	return record, d.DB.Model(&entity.Page1{}).First(&record).Error
}
