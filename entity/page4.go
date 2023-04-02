package entity

import (
	"database/sql/driver"
	"gorm.io/gorm"
	"vizzy/common"
)

type Page4EntitySeries struct {
	Type string `json:"type"`
	Data []int  `json:"data"`
}
type Page4RelEntity struct {
	Title struct {
	} `json:"title"`
	XAxis struct {
		Type     string   `json:"type"`
		Data     []string `json:"data"`
		Show     bool     `json:"show"`
		AxisLine struct {
			Show bool `json:"show"`
		} `json:"axisLine"`
		SplitLine struct {
			Show bool `json:"show"`
		} `json:"splitLine"`
	} `json:"xAxis" gorm:"type:json"`
	YAxis struct {
		Type     string `json:"type"`
		AxisLine struct {
			Show bool `json:"show"`
		} `json:"axisLine"`
		SplitLine struct {
			Show bool `json:"show"`
		} `json:"splitLine"`
	} `json:"yAxis" gorm:"type:json"`
	Series []Page4EntitySeries `json:"series" gorm:"type:json"`
	Grid   struct {
		ContainLabel bool `json:"containLabel"`
		Left         int  `json:"left"`
		Right        int  `json:"right"`
		Top          int  `json:"top"`
		Bottom       int  `json:"bottom"`
	} `json:"grid" gorm:"type:json"`
}
type Page4 struct {
	gorm.Model
	Entity Page4RelEntity `gorm:"type:json"`
}

// Scan Scanner
func (e *Page4RelEntity) Scan(value interface{}) error {
	return common.Scan(&e, value)
}

// Value Valuer
func (e Page4RelEntity) Value() (driver.Value, error) {
	return common.Value(e)
}

func (p *Page4) AfterCreate(tx *gorm.DB) (err error) {
	// 使用事务
	return tx.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err = tx.Model(&Page4{}).Count(&count).Error; err != nil {
			return err
		}
		if count > 1 {
			// 读取旧数据
			var oldData Page4
			if err := tx.First(&oldData).Error; err != nil {
				return err
			}
			// 删除旧数据
			if err := tx.Unscoped().Delete(&oldData).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
