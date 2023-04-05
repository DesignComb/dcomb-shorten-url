package model

import (
	"fmt"
	"main/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type UrlShorten struct {
	ID      uint64 `json:"id" gorm:"primaryKey"`
	Origin  string `json:"origin" gorm:"not null"`
	Short   string `json:"short" gorm:"unique;not null"`
	Clicked uint64 `json:"clickedNum"`
	Random  bool   `json:"isRandom"`
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

	err = db.AutoMigrate(&UrlShorten{})
	if err != nil {
		fmt.Println(err)
	}
}
