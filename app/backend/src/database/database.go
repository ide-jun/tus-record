package database

import (
	"fmt"

	"github.com/ide-jun/tus-record/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetUpDB(host string, port uint16, dbname, username, password string) (err error) {
	// dsn データソースネーム
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		fmt.Sprint(port),
		dbname,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := Migrate(); err != nil {
		return err
	}

	return nil
}

func Migrate() error {
	err := db.AutoMigrate(
		&models.User{},
	)
	return err
}

func GetDB() *gorm.DB {
	return db
}
