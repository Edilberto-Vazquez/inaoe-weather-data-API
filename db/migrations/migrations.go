package migrations

import (
	"fmt"
	"log"
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type Place struct {
	ID   int    `gorm:"<-:create;primaryKey;unique;autoIncrement;not null"`
	Name string `gorm:"<-:create;unique;not null"`
}

type Log struct {
	gorm.Model
	DateTime  string `gorm:"<-:create;not null"`
	Lightning bool   `gorm:"<-:create;not null"`
	Distance  uint8  `gorm:"<-:create;not null"`
	PlaceID   int
	// Place     Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ElectricField struct {
	gorm.Model
	DateTime      time.Time `gorm:"<-:create;not null"`
	ElectricField float32   `gorm:"<-:create;not null"`
	RotorStatus   bool      `gorm:"<-:create;not null"`
	PlaceID       int
	// Place         Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Weathercloud struct {
	gorm.Model
	DateTime time.Time `gorm:"<-:create;not null"`
	TempIn   float32   `gorm:"<-:create;not null"`
	Temp     float32   `gorm:"<-:create;not null"`
	Chill    float32   `gorm:"<-:create;not null"`
	DewIn    float32   `gorm:"<-:create;not null"`
	Dew      float32   `gorm:"<-:create;not null"`
	HeatIn   float32   `gorm:"<-:create;not null"`
	Heat     float32   `gorm:"<-:create;not null"`
	Humin    float32   `gorm:"<-:create;not null"`
	Hum      float32   `gorm:"<-:create;not null"`
	Wspdhi   float32   `gorm:"<-:create;not null"`
	Wspdavg  float32   `gorm:"<-:create;not null"`
	Wdiravg  float32   `gorm:"<-:create;not null"`
	Bar      float32   `gorm:"<-:create;not null"`
	Rain     float32   `gorm:"<-:create;not null"`
	RainRate float32   `gorm:"<-:create;not null"`
	PlaceID  int
	// Place    Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func Migrate() {
	m := gormigrate.New(libs.DBCon, gormigrate.DefaultOptions, []*gormigrate.Migration{})
	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&Place{},
			&Log{},
			&ElectricField{},
			&Weathercloud{},
		)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if err := tx.Exec("ALTER TABLE logs ADD CONSTRAINT fk_logs_place FOREIGN KEY (place_id) REFERENCES places (id)").Error; err != nil {
			fmt.Println(err)
			return err
		}
		if err := tx.Exec("ALTER TABLE electric_fields ADD CONSTRAINT fk_electric_fields_place FOREIGN KEY (place_id) REFERENCES places (id)").Error; err != nil {
			fmt.Println(err)
			return err
		}
		if err := tx.Exec("ALTER TABLE weatherclouds ADD CONSTRAINT fk_weatherclouds_place FOREIGN KEY (place_id) REFERENCES places (id)").Error; err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
