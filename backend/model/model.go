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
	ImageId     uint64 `json:"imageId" gorm:"default:null"`
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

type Image struct {
	ID     uint64 `json:"id" gorm:"primaryKey"`
	UserId uint64 `json:"userId" gorm:"not null"`
	Uri    string `json:"uri" gorm:"unique;not null"`
}

type Platform struct {
	ID      uint64 `json:"id" gorm:"primaryKey"`
	UserId  uint64 `json:"userId" gorm:"default:null"`
	ImageId uint64 `json:"imageId" gorm:"default:null"`
	Name    string `json:"name" gorm:"not null"`
}

type Tree struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	UserId      uint64 `json:"userId" gorm:"default:null"`
	ImageId     uint64 `json:"imageId" gorm:"default:null"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"default:null"`
}

type Link struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	TreeId      uint64 `json:"treeId" gorm:"not null"`
	PlatformId  uint64 `json:"platformId" gorm:"not null"`
	Link        string `json:"link" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"default:null"`
	UserId      uint64 `json:"userId" gorm:"not null"`
	ImageId     uint64 `json:"imageId" gorm:"default:null"`
	Sort        uint64 `json:"sort" gorm:"not null"`
	IsOnlyIcon  bool   `json:"isOnlyIcon" gorm:"default:false;not null"`
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

	err = db.AutoMigrate(&UrlShorten{}, &User{}, &Image{}, &Platform{}, &Tree{}, &Link{})
	if err != nil {
		fmt.Println(err)
	}
}
