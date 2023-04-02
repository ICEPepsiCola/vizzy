package dao

import (
	"gorm.io/gorm"
	"vizzy/entity"
)

type Page5Dao struct {
	DB *gorm.DB
}

func (d Page5Dao) Create(form []entity.Page5) error {
	return d.DB.Transaction(func(tx *gorm.DB) error {
		// 删除整张表的数据
		if err := tx.Unscoped().Where("id >= ?", 1).Delete(&entity.Page5{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 执行批量插入操作
		if err := tx.Create(&form).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})
}

func (d Page5Dao) Query() ([]entity.Page5, error) {
	var record []entity.Page5
	return record, d.DB.Model(&entity.Page5{}).Scan(&record).Error
}
