package connector

import(
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"fmt"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return db
}