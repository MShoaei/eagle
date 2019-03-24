package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // driver
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "host=db user=ms dbname=ms password=asusmhdsh sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
