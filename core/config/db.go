package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase(conn Postgres) (*gorm.DB, error) {

	username := conn.UserName
	password := conn.Password
	host := conn.Host
	port := conn.Port
	database := conn.Database
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
