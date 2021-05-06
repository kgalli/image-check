package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func NewDBConnection() (*Connection, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9970 sslmode=disable TimeZone=Europe/Berlin"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return &Connection{
		DB: db,
	}, err
}

func (c *Connection) Close() error {
	sqlDB, _ := c.DB.DB()
	return sqlDB.Close()
}

func (c *Connection) Migrate(dst ...interface{}) error {
	return c.DB.AutoMigrate(dst...)
}
