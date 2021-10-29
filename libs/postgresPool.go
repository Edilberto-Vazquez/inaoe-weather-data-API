package libs

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
	sqlDB *sql.DB
)

func InitCon() {
	var err error
	dsn := "host=localhost user=edi password=1234 dbname=my_store port=5432 sslmode=disable"
	DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
}

func CloseCon() {
	sqlDB, err := DBCon.DB()
	defer sqlDB.Close()
	if err != nil {
		log.Println(err)
	}
}
