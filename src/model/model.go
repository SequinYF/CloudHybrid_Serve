package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var CHdb *gorm.DB

type UserInfo struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
	Dir    string `json:"dir"`
}

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}


func Init()  {
	CHdb = NewConn()

	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	CHdb.DB().SetMaxIdleConns(10)
	CHdb.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	CHdb.LogMode(true)

	// 自动迁移模式
	CHdb.AutoMigrate(&UserInfo{})
	has := CHdb.HasTable(&UserInfo{})
	if !has {
		CHdb.CreateTable(&UserInfo{})
	}
}
func NewConn() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}
	return db
}
