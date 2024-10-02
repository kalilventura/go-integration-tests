package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func postgresDialector(dns string) gorm.Dialector {
	return postgres.Open(dns)
}

func GormDB(dsn string) *gorm.DB {
	dialector := postgresDialector(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
