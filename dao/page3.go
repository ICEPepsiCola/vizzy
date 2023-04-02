package dao

import (
	"gorm.io/gorm"
	"vizzy/entity"
)

type Page3Dao struct {
	DB *gorm.DB
}

func (d Page3Dao) Create(form entity.Page3) error {
	return d.DB.Model(&entity.Page3{}).Create(&form).Error
}

func (d Page3Dao) Query() (entity.Page3, error) {
	var record entity.Page3
	return record, d.DB.Model(&entity.Page3{}).First(&record).Error
}
