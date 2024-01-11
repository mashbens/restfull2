package util

import (
	"fmt"

	"github.com/mashbens/restfull2/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	POSTGRES DatabaseDriver = "POSTGRES"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	POSTGRES *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch config.Driver {
	case "POSTGRES":
		db.Driver = POSTGRES
		db.POSTGRES = NewPOSTGRES(config)
	default:
		panic("Database driver not supported")
	}
	return &db
}
func NewPOSTGRES(config *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DB_Host,
		config.DB_User,
		config.DB_Pass,
		config.DB_Name,
		config.DB_Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Debug().Msg(dsn)
	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.POSTGRES != nil {
		db, _ := db.POSTGRES.DB()
		db.Close()
	}
}
