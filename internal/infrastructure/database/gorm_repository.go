// infrastructure/database/connection.go
package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func ConnectDB(config Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
