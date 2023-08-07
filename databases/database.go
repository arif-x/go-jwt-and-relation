package databases

import (
	"fmt"

	"gorm.io/driver/postgres"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	// Postgre
	dsn := "host=127.0.0.1 user=postgres password=password dbname=db_name port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// Mysql
	// dsn := "root:@tcp(127.0.0.1:5432)/db_name?charset=utf8mb4&parseTime=True&loc=Local"
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connected to database.")
}
