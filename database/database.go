package database

import "gorm.io/gorm"
import "gorm.io/driver/mysql"
import "../models"

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open( mysql.Open("root:rootroot@/yt_go_auth"), &gorm.Config{})

	if err != nil {
		panic("could not connect ot the database")
	}
	 DB = conn
	conn.AutoMigrate(&models.User{})
}


