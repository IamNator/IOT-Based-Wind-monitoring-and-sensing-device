package storage

import (
	"fmt"
	"log"

	"github.com/IamNator/iot-wind/model"

	"gorm.io/driver/mysql"

	"github.com/IamNator/iot-wind/pkg/environment"

	"gorm.io/gorm"
)

//Storage
//
//type Storage struct {
//	db *gorm.DB
//}
type Storage struct {
	db *gorm.DB
}

//New creates a new storage object ->
//contains the database connection object
func New(env *environment.Env) *Storage {

	db, err := gorm.Open(
		mysql.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=Africa/Lagos",
			env.Get("IOT_PG_ADDRESS"),
			env.Get("IOT_PG_PORT"),
			env.Get("IOT_PG_USER"),
			env.Get("IOT_PG_DATABASE"),
			env.Get("IOT_PG_PASSWORD"),
			env.Get("IOT_PG_SSLMODE"),
		)),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err)
	}

	return &Storage{
		db: db,
	}

}

func Migration(s *Storage) error {
	return s.db.AutoMigrate(&model.Log{})
}