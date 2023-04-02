package entity

import (
	"database/sql/driver"
	"gorm.io/gorm"
	"vizzy/common"
)

type Page1EntitySeries struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Data      []int  `json:"data"`
	LineStyle struct {
		Type string `json:"type"`
	} `json:"lineStyle,omitempty"`
}
type Page1RelEntity struct {
	Title struct {
		Text string `json:"text"`
	} `json:"title"`
	Legend struct {
		Data  []string `json:"data"`
		Right int      `json:"right"`
	} `json:"legend" gorm:"type:json"`
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
	Series []Page1EntitySeries `json:"series" gorm:"type:json"`
	Grid   struct {
		ContainLabel bool `json:"containLabel"`
		Left         int  `json:"left"`
		Right        int  `json:"right"`
		Top          int  `json:"top"`
		Bottom       int  `json:"bottom"`
	} `json:"grid" gorm:"type:json"`
}

// Scan Scanner
func (e *Page1RelEntity) Scan(value interface{}) error {
	return common.Scan(&e, value)
}

// Value Valuer
func (e Page1RelEntity) Value() (driver.Value, error) {
	return common.Value(e)
}

type Page1 struct {
	gorm.Model
	Entity Page1RelEntity `gorm:"type:json"`
}

func (p *Page1) AfterCreate(tx *gorm.DB) (err error) {
	// 使用事务
	return tx.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err = tx.Model(&Page1{}).Count(&count).Error; err != nil {
			return err
		}
		if count > 1 {
			// 读取旧数据
			var oldData Page1
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
