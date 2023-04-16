package model

import (
	"fmt"
	"main/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type UrlShorten struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	UserId      uint64 `json:"userId" gorm:"default:null"`
	Origin      string `json:"origin" gorm:"not null"`
	Title       string `json:"title" gorm:"default:null"`
	Description string `json:"description" gorm:"default:null"`
	Image       string `json:"image" gorm:"default:null"`
	Short       string `json:"short" gorm:"unique;not null"`
	Clicked     uint64 `json:"clickedNum"`
	Random      bool   `json:"isRandom"`
}

type User struct {
	ID                uint64 `json:"id" gorm:"primaryKey"`
	GoogleUserId      string `json:"googleUserId" gorm:"unique;not null"`
	GoogleUserEmail   string `json:"googleUserEmail" gorm:"unique;not null"`
	GoogleUserName    string `json:"googleUserName" gorm:"not null"`
	GoogleUserPicture string `json:"googleUserPicture"`
}

func Setup() {
	dsn := "host=" + config.Val.DbHost +
		" user=" + config.Val.DbUser +
		" password=" + config.Val.DbPwd +
		" dbname=" + config.Val.DbName +
		" port=5432 sslmode=disable TimeZone=Asia/Taipei"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&UrlShorten{},&User{})
	if err != nil {
		fmt.Println(err)
	}
}
