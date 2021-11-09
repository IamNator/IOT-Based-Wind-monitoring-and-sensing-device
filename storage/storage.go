package storage

import (
	"log"

	"github.com/IamNator/iot-wind/model"

	"github.com/IamNator/iot-wind/pkg/environment"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	dsn := env.Get("IOT_MYSQL_DSN")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err != nil {
		log.Fatal(err)
	}

	return &Storage{
		db: db,
	}

}

func Migration(s *Storage) error {
	return s.db.AutoMigrate(&model.Log{}).Error
}
