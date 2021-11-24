package libs

import (
	"database/sql"
	"log"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
	sqlDB *sql.DB
	err   error
)

func InitCon() {
	dsn := "host=" + config.DB_HOST +
		" user=" + config.DB_USER +
		" password=" + config.DB_PASSWORD +
		" dbname=" + config.DB_NAME +
		" port=" + config.DB_PORT +
		" sslmode=disable"
	DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
}

func CloseCon() {
	sqlDB, err = DBCon.DB()
	if err != nil {
		log.Println(err)
	}
	defer sqlDB.Close()
}
