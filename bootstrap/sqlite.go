package bootstrap

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"sync"
	"vizzy/entity"
)

var (
	sqliteOnce sync.Once
	sqliteDB   *gorm.DB
)

func initSqlite() *gorm.DB {
	var err error
	sqliteOnce.Do(func() {
		sqliteDB, err = gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	})
	if err != nil {
		panic("failed to connect database")
	}
	_ = sqliteDB.AutoMigrate(&entity.Page1{})
	_ = sqliteDB.AutoMigrate(&entity.Page2{})
	_ = sqliteDB.AutoMigrate(&entity.Page3{})
	_ = sqliteDB.AutoMigrate(&entity.Page4{})
	_ = sqliteDB.AutoMigrate(&entity.Page5{})
	if err != nil {
		panic("failed to AutoMigrate database")
	}
	return sqliteDB
}
