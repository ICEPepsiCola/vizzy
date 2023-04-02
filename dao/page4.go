package dao

import (
	"gorm.io/gorm"
	"vizzy/entity"
)

type Page4Dao struct {
	DB *gorm.DB
}

func (d Page4Dao) Create(form entity.Page4) error {
	return d.DB.Model(&entity.Page4{}).Create(&form).Error
}

func (d Page4Dao) Query() (entity.Page4, error) {
	var record entity.Page4
	return record, d.DB.Model(&entity.Page4{}).First(&record).Error
}
