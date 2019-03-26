package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // driver
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost user=ms dbname=cc_development password=asusmhdsh sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
