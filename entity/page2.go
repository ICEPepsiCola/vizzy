package entity

import (
	"database/sql/driver"
	"gorm.io/gorm"
	"vizzy/common"
)

type Page2EntitySeries struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Radius string `json:"radius"`
	Data   []struct {
		Value int    `json:"value"`
		Name  string `json:"name"`
	} `json:"data"`
	Emphasis struct {
		ItemStyle struct {
			ShadowBlur    int    `json:"shadowBlur"`
			ShadowOffsetX int    `json:"shadowOffsetX"`
			ShadowColor   string `json:"shadowColor"`
		} `json:"itemStyle"`
	} `json:"emphasis"`
}
type Page2RelEntity struct {
	Title struct {
	} `json:"title"`
	//Legend struct {
	//	Orient string `json:"orient"`
	//	Left   string `json:"left"`
	//} `json:"legend"`
	Series []Page2EntitySeries `json:"series" gorm:"type:json"`
	Grid   struct {
		ContainLabel bool `json:"containLabel"`
		Left         int  `json:"left"`
		Right        int  `json:"right"`
		Top          int  `json:"top"`
		Bottom       int  `json:"bottom"`
	} `json:"grid" gorm:"type:json"`
}
type Page2 struct {
	gorm.Model
	Entity Page2RelEntity `gorm:"type:json"`
}

// Scan Scanner
func (e *Page2RelEntity) Scan(value interface{}) error {
	return common.Scan(&e, value)
}

// Value Valuer
func (e Page2RelEntity) Value() (driver.Value, error) {
	return common.Value(e)
}

func (p *Page2) AfterCreate(tx *gorm.DB) (err error) {
	// 使用事务
	return tx.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err = tx.Model(&Page2{}).Count(&count).Error; err != nil {
			return err
		}
		if count > 1 {
			// 读取旧数据
			var oldData Page2
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
