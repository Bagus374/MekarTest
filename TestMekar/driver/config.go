package driver

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"

	_ "github.com/lib/pq" //library gorm
)

func Connect() *gorm.DB {
	conn := os.Getenv("URL")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
