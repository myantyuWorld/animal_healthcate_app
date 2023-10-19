package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Pets struct {
	Name      string
	BirthDay  time.Time
	Sex       int8 // 0:オス, 1: メス
	CreatedAt time.Time
	UpdatedAt time.Time
}

func seeds(db *gorm.DB) error {
	// TODO: 日本語を投入しようとすると以下警告
	// Error 1366 (HY000): Incorrect string value: '\xE3\x83\xA1\xE3\x82\xB9' for column 'sex' at row 1root@3982857c066f:/go/src/app/seed#
	pet := Pets{Name: "natsu", BirthDay: time.Now(), Sex: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := db.Create(&pet).Error; err != nil {
		fmt.Printf("%+v", err)
	}
	pet1 := Pets{Name: "chai", BirthDay: time.Now(), Sex: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := db.Create(&pet1).Error; err != nil {
		fmt.Printf("%+v", err)
	}
	pet2 := Pets{Name: "mugi", BirthDay: time.Now(), Sex: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := db.Create(&pet2).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	return nil
}

func openConnection() *gorm.DB {
	// 詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
	dsn := "root:hogehoge@tcp(db:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}

func main() {
	// TODO: database.connect.goのConnect()でできるので上記メソッドがいらない
	// db := dbconnect.Connect()
	db := openConnection()
	if err := seeds(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
	defer db.Close()
}
