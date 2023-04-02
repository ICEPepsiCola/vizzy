package dao

import (
	"gorm.io/gorm"
	"vizzy/entity"
)

type Page2Dao struct {
	DB *gorm.DB
}

func (d Page2Dao) Create(form entity.Page2) error {
	return d.DB.Model(&entity.Page2{}).Create(&form).Error
}

func (d Page2Dao) Query() (entity.Page2, error) {
	var record entity.Page2
	return record, d.DB.Model(&entity.Page2{}).First(&record).Error
}
