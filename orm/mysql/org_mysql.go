package main

import (
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Person struct {
	Id int64
	Name string
}

type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

func main() {

	db, error := gorm.Open("mysql", "root:123456@(localhost)/test?charset=utf8&parseTime=True&loc=Local")
	//db, _ := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if error != nil {
		panic(error)
	}

	db.SingularTable(true)

	defer db.Close()

	//db.AutoMigrate(&Test{})
	//db.CreateTable(&Test{})
	db.Create(&Person{
		Name: "123",
	})

	var person Person
	db.First(&person, "id=?", 2)
	println(person.Name)

	//record := db.NewRecord(&Person{
	//	Name: "1222",
	//})

	println()



	/*db.Create(&Like{
		Ip:        "12312",
		Ua:        "123",
		Title:     "123",

		CreatedAt: time.Time{},
	})

	var like Like
	db.First(&like)
	println(like.ID)
	println(like.Ua)*/
}
