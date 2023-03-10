package model

import (
	"fatsharkbot/src/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB

	schemas = []interface{}{
		&Sticker{},
	}
)

func InitDatabase() {
	connectDatabase()
	migrateSchema()
}

func connectDatabase() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.C.Postgresql.Host,
		config.C.Postgresql.User,
		config.C.Postgresql.Password,
		config.C.Postgresql.Dbname,
		config.C.Postgresql.Port,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
}

func migrateSchema() {
	err := db.AutoMigrate(schemas...)
	if err != nil {
		panic(err)
		return
	}
}
